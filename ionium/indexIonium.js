var isFileSystem = require('fs');
var Ionium = /** @class */ (function () {
    function Ionium(parameters) {
        this.$defPath = parameters.defaultPath;
        // Vars for styles
        this.$cssWherePaste = parameters.setPathAndCopyCSS.wherePaste;
        this.$cssNewFolder = parameters.setPathAndCopyCSS.newFolder;
        this.$cssPathForTag = parameters.setPathAndCopyCSS.pathForTag;
        this.$cssWhereCheckPaths = parameters.setPathAndCopyCSS.whereCheckPaths;
        this.$cssMinify = parameters.setPathAndCopyCSS.minify;
        this.$cssHash = parameters.setPathAndCopyCSS.hash;
        this.$tplHTMLS = parameters.templates.searchHTMLS;
        this.$tplTMPLS = parameters.templates.searchTPLS;
        this.$tplDefine = [];
    }
    Ionium.prototype.CopyFileCSS = function () {
        var _this = this;
        this.SetFolders(this.$cssNewFolder);
        var _loop_1 = function (i) {
            var folders = this_1.$cssWhereCheckPaths[i];
            var HTMLfiles = isFileSystem.readdirSync(this_1.$defPath + folders);
            HTMLfiles.forEach(function (item) {
                var isFile = item;
                if (isFile.includes('.html')) {
                    var fileContent = isFileSystem.readFileSync("".concat(_this.$defPath + folders, "/").concat(isFile), 'utf8');
                    var pathToStyle = fileContent.match(/{\$(.*?)\$}/g);
                    if (pathToStyle) {
                        for (i = 0; i < pathToStyle.length; i += 1) {
                            var pathItem = pathToStyle[i];
                            if (pathItem.includes('load_template') || pathItem.includes('define_template')) {
                                // eslint-disable-next-line no-continue
                                continue;
                            }
                            pathToStyle[i] = pathItem.replace(/{\$ path_style="(.*?)" \$}/g, '$1');
                            var fileName = pathToStyle[i].replace(/\/src\/assets\/css\//g, '').split('/');
                            var fileNameString = void 0;
                            if (fileName.length > 1) {
                                for (var f = 0; f < fileName.length - 1; f += 1) {
                                    var fileDir = fileName[f];
                                    if (!isFileSystem.existsSync("".concat(_this.$defPath + _this.$cssWherePaste, "/").concat(fileDir))) {
                                        isFileSystem.mkdirSync("".concat(_this.$defPath + _this.$cssWherePaste, "/").concat(fileDir), function (err) {
                                            if (err)
                                                throw err;
                                        });
                                    }
                                }
                                fileNameString = fileName.join('/');
                            }
                            else {
                                fileNameString = fileName.join('');
                            }
                            isFileSystem.copyFileSync(".".concat(pathToStyle[i]), "".concat(_this.$defPath + _this.$cssWherePaste, "/").concat(fileNameString));
                            if (_this.$cssMinify) {
                                var path = "".concat(_this.$defPath + _this.$cssWherePaste, "/").concat(fileNameString);
                                var fileContentCss = isFileSystem.readFileSync(path, 'utf8');
                                fileContentCss = fileContentCss
                                    .replace(/\n/g, '')
                                    .replace(/\s+/g, ' ');
                                isFileSystem.writeFile(path, fileContentCss, function (err) {
                                    if (err)
                                        throw err;
                                });
                            }
                        }
                    }
                }
            });
            out_i_1 = i;
        };
        var this_1 = this, out_i_1;
        for (var i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
            _loop_1(i);
            i = out_i_1;
        }
    };
    Ionium.prototype.SetPathToCSS = function () {
        var _this = this;
        var _loop_2 = function (i) {
            var folders = this_2.$cssWhereCheckPaths[i];
            var HTMLfiles = isFileSystem.readdirSync(this_2.$defPath + folders);
            HTMLfiles.forEach(function (item) {
                var isFile = item;
                if (isFile.includes('.html')) {
                    var fileContent = isFileSystem.readFileSync("".concat(_this.$defPath + folders, "/").concat(isFile), 'utf8');
                    var pathToStyle = fileContent.match(/{\$(.*?)\$}/g);
                    if (pathToStyle) {
                        for (i = 0; i < pathToStyle.length; i += 1) {
                            var pathItem = pathToStyle[i];
                            if (pathItem.includes('load_template') || pathItem.includes('define_template')) {
                                // eslint-disable-next-line no-continue
                                continue;
                            }
                            pathToStyle[i] = pathItem.replace(/{\$ path_style="(.*?)" \$}/g, '$1');
                            var fileName = pathToStyle[i].replace(/\/src\/assets\/css\//g, '');
                            fileContent = fileContent.replace(pathItem, "<link rel=\"stylesheet\" href=\"".concat(_this.$cssPathForTag + fileName, "\" type=\"text/css\">"));
                        }
                        isFileSystem.writeFile("".concat(_this.$defPath + folders, "/").concat(isFile), fileContent, function (err) {
                            if (err)
                                throw err;
                        });
                    }
                }
            });
            out_i_2 = i;
        };
        var this_2 = this, out_i_2;
        for (var i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
            _loop_2(i);
            i = out_i_2;
        }
    };
    Ionium.prototype.SetFolders = function (path) {
        if (!isFileSystem.existsSync(this.$defPath + path)) {
            isFileSystem.mkdirSync(this.$defPath + path, function (err) {
                if (err)
                    throw err;
            });
        }
    };
    Ionium.prototype.GenerateHash = function () {
        var symbolArr = '1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM';
        var rtsdnr = '';
        for (var i = 0; i < 19; i += 1) {
            var index = Math.floor(Math.random() * symbolArr.length);
            rtsdnr += symbolArr[index];
        }
        return rtsdnr;
    };
    Ionium.prototype.SetHashForFile = function () {
        var _this = this;
        var frequentPaths = [];
        var hashingPaths = [];
        var _loop_3 = function (i) {
            var folders = this_3.$cssWhereCheckPaths[i];
            var HTMLfiles = isFileSystem.readdirSync(this_3.$defPath + folders);
            HTMLfiles.forEach(function (item) {
                var isFile = item;
                if (isFile.includes('.html')) {
                    var fileContent = isFileSystem.readFileSync("".concat(_this.$defPath + folders, "/").concat(isFile), 'utf8');
                    var pathToStyle = fileContent.match(/<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g);
                    if (pathToStyle) {
                        for (var j = 0; j < pathToStyle.length; j += 1) {
                            var pathItem = pathToStyle[j];
                            pathToStyle[j] = pathItem.replace(/<link rel="stylesheet" href="(.*?)" type="text\/css">/g, '$1');
                            var counterFrequent = 0;
                            for (var f = 0; f < frequentPaths.length; f += 1) {
                                var el = frequentPaths[f];
                                if (el === pathToStyle[j]) {
                                    counterFrequent += 1;
                                    break;
                                }
                            }
                            if (counterFrequent === 0) {
                                var paths = pathToStyle[j].split('/');
                                paths[paths.length - 1] = "".concat(_this.GenerateHash(), ".css");
                                frequentPaths.push(pathToStyle[j]);
                                hashingPaths.push(paths.join('/'));
                            }
                            else {
                                counterFrequent = 0;
                            }
                        }
                    }
                }
            });
        };
        var this_3 = this;
        for (var i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
            _loop_3(i);
        }
        var _loop_4 = function (i) {
            var folders = this_4.$cssWhereCheckPaths[i];
            var HTMLfiles = isFileSystem.readdirSync(this_4.$defPath + folders);
            HTMLfiles.forEach(function (item) {
                var isFile = item;
                if (isFile.includes('.html')) {
                    var fileContent = isFileSystem.readFileSync("".concat(_this.$defPath + folders, "/").concat(isFile), 'utf8');
                    var pathToStyle = fileContent.match(/<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g);
                    if (pathToStyle) {
                        for (var j = 0; j < pathToStyle.length; j += 1) {
                            var pathItem = pathToStyle[j];
                            pathToStyle[j] = pathItem.replace(/<link rel="stylesheet" href="(.*?)" type="text\/css">/g, '$1');
                            for (var f = 0; f < frequentPaths.length; f += 1) {
                                var el = frequentPaths[f];
                                if (el === pathToStyle[j]) {
                                    fileContent = fileContent.replace(pathItem, "<link rel=\"stylesheet\" href=\"".concat(hashingPaths[f], "\" type=\"text/css\">"));
                                    var oldStr = frequentPaths[f].substr(1);
                                    var newStr = hashingPaths[f].substr(1);
                                    if (!isFileSystem.existsSync(newStr)) {
                                        isFileSystem.renameSync(oldStr, newStr);
                                    }
                                }
                            }
                        }
                        isFileSystem.writeFile("".concat(_this.$defPath + folders, "/").concat(isFile), fileContent, function (err) {
                            if (err)
                                throw err;
                        });
                    }
                }
            });
        };
        var this_4 = this;
        for (var i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
            _loop_4(i);
        }
    };
    Ionium.prototype.DefineTemplate = function () {
        var _this = this;
        var TMPLfiles = isFileSystem.readdirSync(this.$tplTMPLS);
        TMPLfiles.forEach(function (item) {
            var isFile = item;
            if (isFile.includes(".html")) {
                var fileContent = isFileSystem.readFileSync("".concat(_this.$tplTMPLS, "/").concat(isFile), 'utf8');
                var structureTmpls = {};
                var fileNameContent = fileContent.match(/{\$ define_template="(.*?)" \$}/g);
                structureTmpls.name = fileNameContent[0].replace(/{\$ define_template="(.*?)" \$}/g, '$1');
                structureTmpls.template = fileContent.replace(/{\$ define_template="(.*?)" \$}/g, '')
                    .replace(/{\$ end_template \$}/g, '');
                _this.$tplDefine.push(structureTmpls);
            }
        });
    };
    Ionium.prototype.SetTemplate = function () {
        var _this = this;
        var HTMLfiles = isFileSystem.readdirSync(this.$tplHTMLS);
        HTMLfiles.forEach(function (item) {
            var isFile = item;
            if (isFile.includes(".html")) {
                var fileContent_1 = isFileSystem.readFileSync("".concat(_this.$tplHTMLS, "/").concat(isFile), 'utf8');
                _this.$tplDefine.forEach(function (objTmpl) {
                    fileContent_1 = fileContent_1.replace("{$ load_template=\"".concat(objTmpl.name, "\" $}"), objTmpl.template);
                });
                isFileSystem.writeFile("".concat(_this.$tplHTMLS, "/").concat(isFile), fileContent_1, function (err) {
                    if (err)
                        throw err;
                });
            }
        });
    };
    Ionium.prototype.Start = function () {
        var _this = this;
        setTimeout(function () {
            console.log('Files preparation...');
            _this.CopyFileCSS();
            console.log('Files copied is completed!');
        }, 1000);
        setTimeout(function () {
            console.log('Insert links to files...');
            _this.SetPathToCSS();
            console.log('Inserted links to files is completed!');
            console.log("Start hashing files? Answer: ".concat(_this.$cssHash));
        }, 2000);
        setTimeout(function () {
            _this.DefineTemplate();
        }, 3000);
        setTimeout(function () {
            _this.SetTemplate();
        }, 3500);
        if (this.$cssHash) {
            setTimeout(function () {
                _this.SetHashForFile();
                console.log('Files are hashing is completed!');
            }, 3000);
        }
    };
    return Ionium;
}());
new Ionium({
    defaultPath: './ui',
    setPathAndCopyCSS: {
        newFolder: '/assets/css',
        wherePaste: '/assets/css',
        whereCheckPaths: [
            '/html',
        ],
        pathForTag: '/ui/assets/css/',
        minify: true,
        hash: true
    },
    templates: {
        searchHTMLS: "./ui/html",
        searchTPLS: "./ui/templates/Headers"
    }
}).Start();
