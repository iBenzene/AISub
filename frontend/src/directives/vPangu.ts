import type { Directive } from 'vue'
import pangu from 'pangu'

const vPangu: Directive = {
  mounted(el: HTMLElement) {
    // @ts-ignore: pangu type definition is missing spacingElement
    pangu.spacingElement(el)
  },
  updated(el: HTMLElement) {
    // @ts-ignore: pangu type definition is missing spacingElement
    pangu.spacingElement(el)
  }
}

export default vPangu
