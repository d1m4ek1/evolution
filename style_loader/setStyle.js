let __file_system = require("fs");

const defaultPath = "./ui/assets/";
const HTMLfiles = __file_system.readdirSync("./ui/html/");

if (!__file_system.existsSync(defaultPath + "css")) {
    __file_system.mkdirSync(defaultPath + "css", (err) => {
        if (err) throw err;
    });
}

setTimeout(() => {
    console.log("File preparation...");

    for (const file of HTMLfiles) {
        let fileContent = __file_system.readFileSync(
            "./ui/html/" + file,
            "utf8"
        );
        let pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

        if (pathToStyle) {
            for (let i = 0; i < pathToStyle.length; i++) {
                const pathItem = pathToStyle[i];
                pathToStyle[i] = pathItem.replace(
                    /{\$ path_style="(.*?)" \$}/g,
                    "$1"
                );

                let fileName = pathToStyle[i].replace(
                    /\/src\/assets\/css\//g,
                    ""
                );

                fileName = fileName.split("/");

                if (fileName.length > 1) {
                    for (let j = 0; j < fileName.length - 1; j++) {
                        const fileDir = fileName[j];

                        if (
                            !__file_system.existsSync(
                                defaultPath + "css/" + fileDir
                            )
                        ) {
                            __file_system.mkdirSync(
                                defaultPath + "css/" + fileDir,
                                (err) => {
                                    if (err) throw err;
                                }
                            );
                        }
                    }

                    fileName = fileName.join("/");
                } else {
                    fileName = fileName.join("");
                }

                __file_system.copyFileSync(
                    "." + pathToStyle[i],
                    "./ui/assets/css/" + fileName
                );
            }
        }
    }
}, 1000);

setTimeout(() => {
    console.log("Created directoryes");
}, 2000);

setTimeout(() => {
    for (const file of HTMLfiles) {
        let fileContent = __file_system.readFileSync(
            "./ui/html/" + file,
            "utf8"
        );
        let pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

        if (pathToStyle) {
            for (let i = 0; i < pathToStyle.length; i++) {
                const pathItem = pathToStyle[i];
                pathToStyle[i] = pathItem.replace(
                    /{\$ path_style="(.*?)" \$}/g,
                    "$1"
                );

                let fileName = pathToStyle[i].replace(
                    /\/src\/assets\/css\//g,
                    ""
                );

                fileContent = fileContent.replace(
                    pathItem,
                    `<link rel="stylesheet" href="/ui/assets/css/${fileName}">`
                );
            }

            __file_system.writeFile("./ui/html/" + file, fileContent, (err) => {
                if (err) throw err;
            });
        }
    }

    console.log("Style links written");
}, 3000);
