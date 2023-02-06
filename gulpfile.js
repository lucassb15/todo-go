const {src, dest} = require ("gulp")
const sass = require("gulp-sass")(require("sass"))

exports.default = () => {
    return src("src/**/*.scss")
    .pipe(sass({
        outputStyle: "compressed"
    })).on("error", sass.logError)
    .pipe(dest("assets/dist/css/"));
};