package driver

import (
	"github.com/gobuffalo/packr"
	"github.com/open-policy-agent/opa/ast"

	"github.com/ory/herodot"
	"github.com/ory/x/healthx"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/tracing"

	"github.com/ory/keto/driver/configuration"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
	"github.com/ory/keto/storage"
)

type RegistryBase struct {
	l            *logrusx.Logger
	c            configuration.Provider
	writer       herodot.Writer
	buildVersion string
	buildHash    string
	buildDate    string
	r            Registry
	trc          *tracing.Tracer

	hh *healthx.Handler
	ac *ast.Compiler
	ee *engine.Engine
	le *ladon.Engine
	sh *storage.Handler
}

func (m *RegistryBase) with(r Registry) *RegistryBase {
	m.r = r
	return m
}

func (m *RegistryBase) WithBuildInfo(version, hash, date string) Registry {
	m.buildVersion = version
	m.buildHash = hash
	m.buildDate = date
	return m.r
}
func (m *RegistryBase) BuildVersion() string {
	return m.buildVersion
}

func (m *RegistryBase) BuildDate() string {
	return m.buildDate
}

func (m *RegistryBase) BuildHash() string {
	return m.buildHash
}

func (m *RegistryBase) WithConfig(c configuration.Provider) Registry {
	m.c = c
	return m.r
}

func (m *RegistryBase) Writer() herodot.Writer {
	if m.writer == nil {
		h := herodot.NewJSONWriter(m.Logger())
		h.ErrorEnhancer = nil
		m.writer = h
	}
	return m.writer
}

func (m *RegistryBase) WithLogger(l *logrusx.Logger) Registry {
	m.l = l
	return m.r
}

func (m *RegistryBase) Logger() *logrusx.Logger {
	if m.l == nil {
		m.l = logrusx.New("ORY Keto", m.buildVersion)
	}
	return m.l
}

func (m *RegistryBase) EngineCompiler() *ast.Compiler {
	if m.ac == nil {
		box := packr.NewBox("../engine/ladon/rego")
		compiler, err := engine.NewCompiler(box, m.Logger())
		if err != nil {
			m.Logger().WithError(err).Fatalf("Unable to initialize compiler")
		}
		m.ac = compiler
	}
	return m.ac
}

func (m *RegistryBase) Engine() *engine.Engine {
	if m.ee == nil {
		m.ee = engine.NewEngine(m.EngineCompiler(), m.Writer())
	}
	return m.ee
}

func (m *RegistryBase) LadonEngine() *ladon.Engine {
	if m.le == nil {
		m.le = ladon.NewEngine(m.r.StorageManager(), m.StorageHandler(), m.Engine(), m.Writer())
	}
	return m.le
}

func (m *RegistryBase) StorageHandler() *storage.Handler {
	if m.sh == nil {
		m.sh = storage.NewHandler(m.r.StorageManager(), m.Writer())
	}
	return m.sh
}

func (m *RegistryBase) HealthHandler() *healthx.Handler {
	if m.hh == nil {
		m.hh = healthx.NewHandler(m.Writer(), m.buildVersion, healthx.ReadyCheckers{
			"database": m.r.Ping,
		})
	}

	return m.hh
}

func (m *RegistryBase) Tracer() *tracing.Tracer {
	if m.trc == nil {
		m.trc = &tracing.Tracer{
			ServiceName:  m.c.TracingServiceName(),
			JaegerConfig: m.c.TracingJaegerConfig(),
			Provider:     m.c.TracingProvider(),
			Logger:       m.Logger(),
		}

		if err := m.trc.Setup(); err != nil {
			m.Logger().WithError(err).Fatalf("Unable to initialize Tracer.")
		}
	}

	return m.trc
}
