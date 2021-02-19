it("works", async () => {
    console.log = (...args) => {
        expect(args.length).toBe(1)
        expect(args[0]).toBe("Allowed")
    }

    await import("./index.js")
})
