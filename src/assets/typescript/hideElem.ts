import {DirectiveBinding} from "vue/types/options";
import {VNode} from "vue";

export default {
  update(el: Element, binding: DirectiveBinding, vnode: VNode) {
    const binVal = binding.value;
    const vnodeContext = vnode.context;
    const elem = binVal.id !== undefined ? document.getElementById(binVal.id) : el;

    document.addEventListener('click', (event: MouseEvent) => {
      const element: HTMLInputElement = <HTMLInputElement>event.target
      if (binVal.val.length === 3) {
        Object.keys(vnodeContext).forEach((item: string) => {
          const elemObject:string = item;

          if (elemObject === binVal.val[0]) {
            if (binVal.val[1] === binVal.id && vnodeContext[elemObject][binVal.val[1]][binVal.val[2]]
                && !el.contains(element)) {
              if (binVal.id !== undefined && !elem.contains(element)) {
                vnodeContext[elemObject][binVal.val[1]][binVal.val[2]] = false;
              } else if (binVal.id === undefined) {
                vnodeContext[elemObject][binVal.val[1]][binVal.val[2]] = false;
              }
            }
          }
        });
      }

      if (binVal.val.length === 2) {
        Object.keys(vnodeContext).forEach((item:string) => {
          const elemObject: string = item;

          if (elemObject === binVal.val[0]) {
            if (vnodeContext[elemObject][binVal.val[1]] && !el.contains(element)) {
              if (binVal.id !== undefined && !elem.contains(element)) {
                vnodeContext[elemObject][binVal.val[1]] = false;
              } else if (binVal.id === undefined) {
                vnodeContext[elemObject][binVal.val[1]] = false;
              }
            }
          }
        });
      }

      if (binVal.val.length === 1) {
        Object.keys(vnodeContext).forEach((item: string) => {
          const elemObject: string = item;

          if (elemObject === binVal.val[0]) {
            if (vnodeContext[elemObject] && !el.contains(element)) {
              if (binVal.id !== undefined && !elem.contains(element)) {
                vnodeContext[elemObject] = false;
              } else if (binVal.id === undefined) {
                vnodeContext[elemObject] = false;
              }
            }
          }
        });
      }
    });
  },
};
