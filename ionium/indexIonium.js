const isFileSystem = require('fs');

class Ionium {
  constructor(
    parameters = {
      defaultPath: String,
      setPathAndCopyCSS: {
        newFolder: String,
        wherePaste: String,
        whereCheckPaths: Array,
        pathForTag: String,
        minify: Boolean,
        hash: Boolean,
      },
      templates: {
        searchHTMLS: String,
        searchTPLS: String,
      },
    },
  ) {
    this.$defPath = String(parameters.defaultPath);
    // Vars for styles
    this.$cssWherePaste = String(parameters.setPathAndCopyCSS.wherePaste);
    this.$cssNewFolder = String(parameters.setPathAndCopyCSS.newFolder);
    this.$cssPathForTag = String(parameters.setPathAndCopyCSS.pathForTag);
    this.$cssWhereCheckPaths = parameters.setPathAndCopyCSS.whereCheckPaths;
    this.$cssMinify = Boolean(parameters.setPathAndCopyCSS.minify);
    this.$cssHash = parameters.setPathAndCopyCSS.hash || false;
  }

  #CopyFileCSS() {
    this.#SetFolders(this.$cssNewFolder);

    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders = this.$cssWhereCheckPaths[i];
      const HTMLfiles = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item = String) => {
        const isFile = item;

        if (isFile.includes('.html')) {
          const fileContent = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

          if (pathToStyle) {
            for (i = 0; i < pathToStyle.length; i += 1) {
              const pathItem = pathToStyle[i];
              pathToStyle[i] = pathItem.replace(
                /{\$ path_style="(.*?)" \$}/g,
                '$1',
              );

              let fileName = pathToStyle[i].replace(
                /\/src\/assets\/css\//g,
                '',
              );

              fileName = fileName.split('/');

              if (fileName.length > 1) {
                for (let f = 0; f < fileName.length - 1; f += 1) {
                  const fileDir = fileName[f];

                  if (
                    !isFileSystem.existsSync(
                      `${this.$defPath + this.$cssWherePaste}/${fileDir}`,
                    )
                  ) {
                    isFileSystem.mkdirSync(
                      `${this.$defPath + this.$cssWherePaste}/${fileDir}`,
                      (err) => {
                        if (err) throw err;
                      },
                    );
                  }
                }

                fileName = fileName.join('/');
              } else {
                fileName = fileName.join('');
              }

              isFileSystem.copyFileSync(
                `.${pathToStyle[i]}`,
                `${this.$defPath + this.$cssWherePaste}/${fileName}`,
              );

              if (this.$cssMinify) {
                const path = `${
                  this.$defPath + this.$cssWherePaste
                }/${fileName}`;

                let fileContentCss = isFileSystem.readFileSync(path, 'utf8');

                fileContentCss = fileContentCss
                  .replace(/\n/g, '')
                  .replace(/\s+/g, ' ');
                isFileSystem.writeFile(path, fileContentCss, (err) => {
                  if (err) throw err;
                });
              }
            }
          }
        }
      });
    }
  }

  #SetPathToCSS() {
    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders = this.$cssWhereCheckPaths[i];
      const HTMLfiles = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item = String) => {
        const isFile = item;

        if (isFile.includes('.html')) {
          let fileContent = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

          if (pathToStyle) {
            for (i = 0; i < pathToStyle.length; i += 1) {
              const pathItem = pathToStyle[i];
              pathToStyle[i] = pathItem.replace(
                /{\$ path_style="(.*?)" \$}/g,
                '$1',
              );

              const fileName = pathToStyle[i].replace(
                /\/src\/assets\/css\//g,
                '',
              );

              fileContent = fileContent.replace(
                pathItem,
                `<link rel="stylesheet" href="${
                  this.$cssPathForTag + fileName
                }" type="text/css">`,
              );
            }

            isFileSystem.writeFile(
              `${this.$defPath + folders}/${isFile}`,
              fileContent,
              (err) => {
                if (err) throw err;
              },
            );
          }
        }
      });
    }
  }

  #SetFolders(path) {
    if (!isFileSystem.existsSync(this.$defPath + path)) {
      isFileSystem.mkdirSync(this.$defPath + path, (err = Object) => {
        if (err) throw err;
      });
    }
  }

  #GenerateHash() {
    const symbolArr = '1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM';
    let rtsdnr = '';
    for (let i = 0; i < 19; i += 1) {
      const index = Math.floor(Math.random() * symbolArr.length);
      rtsdnr += symbolArr[index];
    }
    return rtsdnr;
  }

  #SetHashForFile() {
    const frequentPaths = [];
    const hashingPaths = [];

    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders = this.$cssWhereCheckPaths[i];
      const HTMLfiles = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item = String) => {
        const isFile = item;

        if (isFile.includes('.html')) {
          const fileContent = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle = fileContent.match(
            /<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g,
          );

          if (pathToStyle) {
            for (let j = 0; j < pathToStyle.length; j += 1) {
              const pathItem = pathToStyle[j];
              pathToStyle[j] = pathItem.replace(
                /<link rel="stylesheet" href="(.*?)" type="text\/css">/g,
                '$1',
              );

              let counterFrequent = 0;

              for (let f = 0; f < frequentPaths.length; f += 1) {
                const el = frequentPaths[f];

                if (el === pathToStyle[j]) {
                  counterFrequent += 1;
                  break;
                }
              }

              if (counterFrequent === 0) {
                const paths = pathToStyle[j].split('/');
                paths[paths.length - 1] = `${this.#GenerateHash()}.css`;

                frequentPaths.push(pathToStyle[j]);
                hashingPaths.push(paths.join('/'));
              } else {
                counterFrequent = 0;
              }
            }
          }
        }
      });
    }

    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders = this.$cssWhereCheckPaths[i];
      const HTMLfiles = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item = String) => {
        const isFile = item;

        if (isFile.includes('.html')) {
          let fileContent = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle = fileContent.match(
            /<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g,
          );

          if (pathToStyle) {
            for (let j = 0; j < pathToStyle.length; j += 1) {
              const pathItem = pathToStyle[j];
              pathToStyle[j] = pathItem.replace(
                /<link rel="stylesheet" href="(.*?)" type="text\/css">/g,
                '$1',
              );

              for (let f = 0; f < frequentPaths.length; f += 1) {
                const el = frequentPaths[f];

                if (el === pathToStyle[j]) {
                  fileContent = fileContent.replace(
                    pathItem,
                    `<link rel="stylesheet" href="${hashingPaths[f]}" type="text/css">`,
                  );
                  const oldStr = frequentPaths[f].substr(1);
                  const newStr = hashingPaths[f].substr(1);

                  if (!isFileSystem.existsSync(newStr)) {
                    isFileSystem.renameSync(oldStr, newStr);
                  }
                }
              }
            }

            isFileSystem.writeFile(
              `${this.$defPath + folders}/${isFile}`,
              fileContent,
              (err) => {
                if (err) throw err;
              },
            );
          }
        }
      });
    }
  }

  Start() {
    setTimeout(() => {
      console.log('Files preparation...');
      this.#CopyFileCSS();
      console.log('Files copied is completed!');
    }, 1000);

    setTimeout(() => {
      console.log('Insert links to files...');
      this.#SetPathToCSS();
      console.log('Inserted links to files is completed!');
      console.log(`Start hashing files? Answer: ${this.$cssHash}`);
    }, 2000);

    if (this.$cssHash) {
      setTimeout(() => {
        this.#SetHashForFile();
        console.log('Files are hashing is completed!');
      }, 3000);
    }
  }
}

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
    hash: true,
  },
}).Start();
