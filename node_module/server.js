let express = require("express")
let app = express()
let axios = require("axios")

app.get("/", async (req, res, next) => {
    try {
        let response = await axios.get("http://localhost:8080/users")
        console.log(response.data)
    } catch (err) {
        console.error(err)
    }
    next()
})

app.listen(4000, () => {
    console.log("Server nodeJS listening on port 4000 ...")
})