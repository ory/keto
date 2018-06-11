/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package legacy

import "github.com/rubenv/sql-migrate"

var HydraLegacyMigrations = map[string]*migrate.MemoryMigrationSource{
	"postgres": {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`ALTER TABLE hydra_warden_group RENAME TO keto_role`,
					`ALTER TABLE hydra_warden_group_member RENAME COLUMN group_id TO role_id`,
					`ALTER TABLE hydra_warden_group_member RENAME TO keto_role_member`,
					`ALTER TABLE hydra_policy_migration RENAME TO keto_policy_migration`,
					`ALTER TABLE hydra_groups_migration RENAME TO keto_role_migration`,
				},
				Down: []string{
					`ALTER TABLE keto_role RENAME TO hydra_warden_group`,
					`ALTER TABLE hydra_warden_group_member RENAME COLUMN role_id TO group_id`,
					`ALTER TABLE keto_role_member RENAME TO hydra_warden_group_member`,
					`ALTER TABLE keto_policy_migration RENAME TO hydra_policy_migration`,
					`ALTER TABLE keto_role_migration RENAME TO hydra_groups_migration`,
				},
			},
		},
	},
	"mysql": {
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{
					`RENAME TABLE hydra_warden_group RENAME TO keto_role`,
					`ALTER TABLE hydra_warden_group_member CHANGE group_id role_id varchar(255)`,
					`RENAME TABLE hydra_warden_group_member RENAME TO keto_role_member`,
					`RENAME TABLE hydra_policy_migration RENAME TO keto_policy_migration`,
					`RENAME TABLE hydra_groups_migration RENAME TO keto_role_migration`,
				},
				Down: []string{
					`RENAME TABLE keto_role TO hydra_warden_group`,
					`ALTER TABLE hydra_warden_group_member CHANGE role_id group_id varchar(255)`,
					`RENAME TABLE keto_role_member TO hydra_warden_group_member`,
					`RENAME TABLE keto_policy_migration TO hydra_policy_migration`,
					`RENAME TABLE keto_role_migration TO hydra_groups_migration`,
				},
			},
		},
	},
}
