import {
  InputText,
  ComboSimple,
  Button,
  RadioButton,
  SelectTableList,
  SidePanel,
  Card,
  FormGroupInput,
  ModalDialog,
  DropDown,
  Accordion,
  DropMenu,
  HorizontalTabs,
  VerticalTabs,
  DavinciTable,
  SessionWatch,
  FindBoxList,
  InputCheck,
  SwitchBox
} from "@/components/index";

import { Drag, Drop } from "vue-drag-drop";

const GlobalComponents = {
  install(Vue) {
    Vue.component("drag", Drag);
    Vue.component("drop", Drop);
    Vue.component("accordion", Accordion);

    Vue.component("davinci-table", DavinciTable);
    Vue.component("select-table", SelectTableList);
    Vue.component("side-panel", SidePanel);
    Vue.component("combo-simple", ComboSimple);
    Vue.component("find-box", FindBoxList);
    Vue.component("h-tabs", HorizontalTabs);
    Vue.component("v-tabs", VerticalTabs);
    Vue.component("m-dialog", ModalDialog);
    Vue.component("fg-input", FormGroupInput);
    Vue.component("input-text", InputText);
    Vue.component("check-box", InputCheck);
    Vue.component("switch-box", SwitchBox);

    Vue.component("session-watch", SessionWatch);

    Vue.component("drop-down", DropDown);
    Vue.component("drop-menu", DropMenu);

    Vue.component("card", Card);
    Vue.component("d-button", Button);
    Vue.component("d-radio", RadioButton);
  }
};

export default GlobalComponents;
