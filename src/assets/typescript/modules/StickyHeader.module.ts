const MODULE_STICKY_HEADER = (): void => {
    const headerStickyElement: HTMLElement = <HTMLElement>document.querySelector('.header_sticky')
    const mainElement: HTMLElement = <HTMLElement>document.querySelector('.main')

    headerStickyElement.style.transform = 'translateX(-200px)';

    mainElement.classList.remove('main_squeeze_before_add');
    mainElement.classList.add('main_squeeze_before_remove');

    setTimeout(() => {
        mainElement.classList.remove('main_squeeze_before_remove');
    }, 490);
}

export default MODULE_STICKY_HEADER;
