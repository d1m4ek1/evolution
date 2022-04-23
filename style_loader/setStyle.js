const isFileSystem = require('fs');

const defaultPath = './ui/assets/';
const HTMLfiles = isFileSystem.readdirSync('./ui/html/');

if (!isFileSystem.existsSync(`${defaultPath}css`)) {
  isFileSystem.mkdirSync(`${defaultPath}css`, (err) => {
    if (err) throw err;
  });
}

setTimeout(() => {
  console.log('File preparation...');

  HTMLfiles.forEach((item = String) => {
    const isFile = item;
    const fileContent = isFileSystem.readFileSync(
      `./ui/html/${isFile}`,
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
                `${defaultPath}css/${fileDir}`,
              )
            ) {
              isFileSystem.mkdirSync(
                `${defaultPath}css/${fileDir}`,
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
          `./ui/assets/css/${fileName}`,
        );
      }
    }
  });
}, 1000);

setTimeout(() => {
  console.log('Created directoryes');
}, 2000);

setTimeout(() => {
  HTMLfiles.forEach((item = String) => {
    const isFile = item;
    let fileContent = isFileSystem.readFileSync(
      `./ui/html/${isFile}`,
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
          `<link rel="stylesheet" href="/ui/assets/css/${fileName}">`,
        );
      }

      isFileSystem.writeFile(`./ui/html/${isFile}`, fileContent, (err) => {
        if (err) throw err;
      });
    }
  });

  console.log('Style links written');
}, 3000);
