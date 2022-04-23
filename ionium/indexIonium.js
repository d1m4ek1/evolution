const isFileSystem = require('fs');

class Ionium {
  constructor(parameters = {
    defaultPath: String,
    setPathAndCopyCSS: {
      newFolder: String,
      whereCopy: String,
      whereCheckPaths: String,
      pathForTag: String,
      minify: Boolean,
    },
    templates: {
      searchHTMLS: String,
      searchTPLS: String,
    },
  }) {
    this.$defPath = String(parameters.defaultPath);
    // Vars for styles
    this.$cssWhereCopy = String(parameters.setPathAndCopyCSS.whereCopy);
    this.$cssNewFolder = String(parameters.setPathAndCopyCSS.newFolder);
    this.$cssPathForTag = String(parameters.setPathAndCopyCSS.pathForTag);
    this.$cssWhereCheckPaths = String(parameters.setPathAndCopyCSS.whereCheckPaths);
    this.$cssMinify = Boolean(parameters.setPathAndCopyCSS.minify);
  }

  #CopyFileCSS() {
    this.#SetFolders(this.$cssNewFolder);

    const HTMLfiles = isFileSystem.readdirSync(this.$defPath + this.$cssWhereCheckPaths);

    HTMLfiles.forEach((item = String) => {
      const isFile = item;
      const fileContent = isFileSystem.readFileSync(
        `${this.$defPath + this.$cssWhereCheckPaths}/${isFile}`,
        'utf8',
      );
      const pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

      if (pathToStyle) {
        for (let i = 0; i < pathToStyle.length; i += 1) {
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
            for (let j = 0; j < fileName.length - 1; j += 1) {
              const fileDir = fileName[j];

              if (
                !isFileSystem.existsSync(
                  `${this.$defPath + this.$cssWhereCopy}/${fileDir}`,
                )
              ) {
                isFileSystem.mkdirSync(
                  `${this.$defPath + this.$cssWhereCopy}/${fileDir}`,
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
            `${this.$defPath + this.$cssWhereCopy}/${fileName}`,
          );

          if (this.$cssMinify) {
            const path = `${this.$defPath + this.$cssWhereCopy}/${fileName}`;

            let fileContentCss = isFileSystem.readFileSync(path, 'utf8');

            fileContentCss = fileContentCss.replace(/\n/g, '').replace(/\s+/g, ' ');
            isFileSystem.writeFile(path, fileContentCss, (err) => {
              if (err) throw err;
            });
          }
        }
      }
    });
  }

  #SetPathToCSS() {
    const HTMLfiles = isFileSystem.readdirSync(this.$defPath + this.$cssWhereCheckPaths);

    HTMLfiles.forEach((item = String) => {
      const isFile = item;
      let fileContent = isFileSystem.readFileSync(
        `${this.$defPath + this.$cssWhereCheckPaths}/${isFile}`,
        'utf8',
      );
      const pathToStyle = fileContent.match(/{\$(.*?)\$}/g);

      if (pathToStyle) {
        for (let i = 0; i < pathToStyle.length; i += 1) {
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
            `<link rel="stylesheet" href="${this.$cssPathForTag + fileName}">`,
          );
        }

        isFileSystem.writeFile(`${this.$defPath
          + this.$cssWhereCheckPaths}/${isFile}`, fileContent, (err) => {
          if (err) throw err;
        });
      }
    });
  }

  #SetFolders(path) {
    if (!isFileSystem.existsSync(this.$defPath + path)) {
      isFileSystem.mkdirSync(this.$defPath + path, (err = Object) => {
        if (err) throw err;
      });
    }
  }

  Start() {
    setTimeout(() => {
      this.#CopyFileCSS();

      console.log('File preparation...');
    }, 1000);

    setTimeout(() => {
      this.#SetPathToCSS();
      console.log('Style links written');
    }, 2000);
  }
}

new Ionium({
  defaultPath: './ui',
  setPathAndCopyCSS: {
    newFolder: '/assets/css',
    whereCopy: '/assets/css',
    whereCheckPaths: '/html',
    pathForTag: '/ui/assets/css/',
    minify: true,
  },
}).Start();
