import FormGroupInput from "./UserInterface/Inputs/formGroupInput.vue";
import ComboSimple from "./UserInterface/Inputs/ComboSimple.vue";
import InputCheck from "./UserInterface/Inputs/InputCheck.vue";
import SwitchBox from "./UserInterface/Inputs/SwitchBox.vue";
import InputText from "./UserInterface/Inputs/InputText.vue";
import FindBoxList from "./UserInterface/Inputs/FindBoxList.vue";
import SelectTableList from "./UserInterface/Inputs/SelectTableList.vue";
import LabelList from "./UserInterface/Inputs/LabelList.vue";

import ModalDialog from "./UserInterface/Alert/ModalDialog.vue";

import DropDown from "./UserInterface/Menus/Dropdown.vue";
import DropMenu from "./UserInterface/Menus/DropMenu.vue";

import Button from "./UserInterface/Button/Button";
import RadioButton from "./UserInterface/Button/RadioButton";

import Card from "./UserInterface/Cards/Card.vue";
import ChartCard from "./UserInterface/Cards/ChartCard.vue";
import StatsCard from "./UserInterface/Cards/StatsCard.vue";

import SidebarPlugin from "./UserInterface/SidebarPlugin/index";

import HorizontalTabs from "./UserInterface/Containers/HorizontalTabs.vue";
import VerticalTabs from "./UserInterface/Containers/VerticalTabs.vue";
import DavinciTable from "./UserInterface/Containers/DavinciTable.vue";
import Pagination from "./UserInterface/Containers/Pagination.vue";
import SidePanel from "./UserInterface/Containers/SidePanel.vue";
import Accordion from "./UserInterface/Containers/Accordion.vue";

import SessionWatch from "./Utils/SessionWatch.vue";

let components = {
  ComboSimple,
  LabelList,
  Pagination,
  SidePanel,
  Accordion,
  HorizontalTabs,
  FindBoxList,
  VerticalTabs,
  SelectTableList,
  FormGroupInput,
  ModalDialog,
  InputCheck,
  DavinciTable,
  SwitchBox,
  Card,
  ChartCard,
  StatsCard,
  InputText,
  DropDown,
  DropMenu,
  SessionWatch,
  Button,
  RadioButton,
  SidebarPlugin
};

export default components;

export {
  ComboSimple,
  LabelList,
  HorizontalTabs,
  VerticalTabs,
  DavinciTable,
  SidePanel,
  Accordion,
  FormGroupInput,
  FindBoxList,
  InputText,
  SelectTableList,
  InputCheck,
  SwitchBox,
  ModalDialog,
  Card,
  Pagination,
  ChartCard,
  StatsCard,
  DropDown,
  DropMenu,
  SessionWatch,
  Button,
  RadioButton,
  SidebarPlugin
};
