const isFileSystem = require('fs');

interface iParamIonium {
  defaultPath: string,
  setPathAndCopyCSS: {
    newFolder: string,
    wherePaste: string,
    whereCheckPaths: Array<string>,
    pathForTag: string,
    minify: boolean,
    hash: boolean,
  },
  templates: {
    searchHTMLS: string,
    searchTPLS: string,
  },
}

interface iStructureTmpls {
  name: string,
  template: string,
}

class Ionium {
  protected $defPath: string
  protected $cssWherePaste: string
  protected $cssNewFolder: string
  protected $cssPathForTag: string
  protected $cssWhereCheckPaths: Array<string>
  protected $cssMinify: boolean
  protected $cssHash: boolean
  protected $tplHTMLS: string
  protected $tplTMPLS: string
  protected $tplDefine: Array<Partial<iStructureTmpls>>


  constructor(parameters: iParamIonium) {
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

  private CopyFileCSS() {
    this.SetFolders(this.$cssNewFolder);

    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders: string = this.$cssWhereCheckPaths[i];
      const HTMLfiles: Array<string> = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item: string) => {
        const isFile: string = item;

        if (isFile.includes('.html')) {
          const fileContent: string = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle: Array<string> = fileContent.match(/{\$(.*?)\$}/g);

          if (pathToStyle) {
            for (i = 0; i < pathToStyle.length; i += 1) {
              const pathItem: string = pathToStyle[i];

              if (pathItem.includes('load_template') || pathItem.includes('define_template')) {
                // eslint-disable-next-line no-continue
                continue;
              }

              pathToStyle[i] = pathItem.replace(
                /{\$ path_style="(.*?)" \$}/g,
                '$1',
              );

              let fileName: Array<string> = pathToStyle[i].replace(
                /\/src\/assets\/css\//g,
                '',
              ).split('/');

              let fileNameString: string

              if (fileName.length > 1) {
                for (let f = 0; f < fileName.length - 1; f += 1) {
                  const fileDir: string = fileName[f];

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

                fileNameString = fileName.join('/');
              } else {
                fileNameString = fileName.join('');
              }

              isFileSystem.copyFileSync(
                `.${pathToStyle[i]}`,
                `${this.$defPath + this.$cssWherePaste}/${fileNameString}`,
              );

              if (this.$cssMinify) {
                const path: string = `${
                  this.$defPath + this.$cssWherePaste
                }/${fileNameString}`;

                let fileContentCss: string = isFileSystem.readFileSync(path, 'utf8');

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

  private SetPathToCSS() {
    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders: string = this.$cssWhereCheckPaths[i];
      const HTMLfiles: Array<string> = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item: string) => {
        const isFile: string = item;

        if (isFile.includes('.html')) {
          let fileContent = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle: Array<string> = fileContent.match(/{\$(.*?)\$}/g);

          if (pathToStyle) {
            for (i = 0; i < pathToStyle.length; i += 1) {
              const pathItem: string = pathToStyle[i];

              if (pathItem.includes('load_template') || pathItem.includes('define_template')) {
                // eslint-disable-next-line no-continue
                continue;
              }

              pathToStyle[i] = pathItem.replace(
                /{\$ path_style="(.*?)" \$}/g,
                '$1',
              );

              const fileName: string = pathToStyle[i].replace(
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

  private SetFolders(path) {
    if (!isFileSystem.existsSync(this.$defPath + path)) {
      isFileSystem.mkdirSync(this.$defPath + path, (err: Object) => {
        if (err) throw err;
      });
    }
  }

  private GenerateHash() {
    const symbolArr: string = '1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM';
    let rtsdnr: string = '';

    for (let i = 0; i < 19; i += 1) {
      const index: number = Math.floor(Math.random() * symbolArr.length);
      rtsdnr += symbolArr[index];
    }

    return rtsdnr;
  }

  private SetHashForFile() {
    const frequentPaths: Array<string> = [];
    const hashingPaths: Array<string> = [];

    for (let i = 0; i < this.$cssWhereCheckPaths.length; i += 1) {
      const folders: string = this.$cssWhereCheckPaths[i];
      const HTMLfiles: Array<string> = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item: string) => {
        const isFile: string = item;

        if (isFile.includes('.html')) {
          const fileContent: string = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle: Array<string> = fileContent.match(
            /<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g,
          );

          if (pathToStyle) {
            for (let j = 0; j < pathToStyle.length; j += 1) {
              const pathItem: string = pathToStyle[j];
              pathToStyle[j] = pathItem.replace(
                /<link rel="stylesheet" href="(.*?)" type="text\/css">/g,
                '$1',
              );

              let counterFrequent: number = 0;

              for (let f = 0; f < frequentPaths.length; f += 1) {
                const el: string = frequentPaths[f];

                if (el === pathToStyle[j]) {
                  counterFrequent += 1;
                  break;
                }
              }

              if (counterFrequent === 0) {
                const paths: Array<string> = pathToStyle[j].split('/');
                paths[paths.length - 1] = `${this.GenerateHash()}.css`;

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
      const folders: string = this.$cssWhereCheckPaths[i];
      const HTMLfiles: Array<string> = isFileSystem.readdirSync(this.$defPath + folders);

      HTMLfiles.forEach((item: string) => {
        const isFile: string = item;

        if (isFile.includes('.html')) {
          let fileContent: string = isFileSystem.readFileSync(
            `${this.$defPath + folders}/${isFile}`,
            'utf8',
          );
          const pathToStyle: Array<string> = fileContent.match(
            /<link rel="stylesheet" href="\/ui\/assets\/css\/(.*?)" type="text\/css">/g,
          );

          if (pathToStyle) {
            for (let j = 0; j < pathToStyle.length; j += 1) {
              const pathItem: string = pathToStyle[j];
              pathToStyle[j] = pathItem.replace(
                /<link rel="stylesheet" href="(.*?)" type="text\/css">/g,
                '$1',
              );

              for (let f = 0; f < frequentPaths.length; f += 1) {
                const el: string = frequentPaths[f];

                if (el === pathToStyle[j]) {
                  fileContent = fileContent.replace(
                    pathItem,
                    `<link rel="stylesheet" href="${hashingPaths[f]}" type="text/css">`,
                  );
                  const oldStr: string = frequentPaths[f].substr(1);
                  const newStr: string = hashingPaths[f].substr(1);

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

  private DefineTemplate() {
    const TMPLfiles: Array<string> = isFileSystem.readdirSync(this.$tplTMPLS);

    TMPLfiles.forEach((item: string) => {
      const isFile: string = item;

      if (isFile.includes(".html")) {
        const fileContent: string = isFileSystem.readFileSync(
          `${this.$tplTMPLS}/${isFile}`,
          'utf8',
        );

        let structureTmpls: Partial<iStructureTmpls> = {}

        const fileNameContent: Array<string> = fileContent.match(/{\$ define_template="(.*?)" \$}/g);
        structureTmpls.name = fileNameContent[0].replace(/{\$ define_template="(.*?)" \$}/g, '$1');

        structureTmpls.template = fileContent.replace(/{\$ define_template="(.*?)" \$}/g, '')
          .replace(/{\$ end_template \$}/g, '');

        this.$tplDefine.push(structureTmpls);
      }
    });
  }

  private SetTemplate() {
    const HTMLfiles: Array<string> = isFileSystem.readdirSync(this.$tplHTMLS);

    HTMLfiles.forEach((item: string) => {
      const isFile: string = item;

      if (isFile.includes(".html")) {
        let fileContent: string = isFileSystem.readFileSync(
          `${this.$tplHTMLS}/${isFile}`,
          'utf8',
        );

        this.$tplDefine.forEach((objTmpl: Partial<iStructureTmpls>) => {
          fileContent = fileContent.replace(`{$ load_template="${objTmpl.name}" $}`, objTmpl.template);
        });

        isFileSystem.writeFile(
          `${this.$tplHTMLS}/${isFile}`,
          fileContent,
          (err) => {
            if (err) throw err;
          },
        );
      }
    });
  }

  Start() {
    setTimeout(() => {
      console.log('Files preparation...');
      this.CopyFileCSS();
      console.log('Files copied is completed!');
    }, 1000);

    setTimeout(() => {
      console.log('Insert links to files...');
      this.SetPathToCSS();
      console.log('Inserted links to files is completed!');
      console.log(`Start hashing files? Answer: ${this.$cssHash}`);
    }, 2000);

    setTimeout(() => {
      this.DefineTemplate();
    }, 3000);

    setTimeout(() => {
      this.SetTemplate();
    }, 3500);

    if (this.$cssHash) {
      setTimeout(() => {
        this.SetHashForFile();
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
  templates: {
    searchHTMLS: "./ui/html",
    searchTPLS: "./ui/templates/Headers"
  },
}).Start();
