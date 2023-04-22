import { defineStore } from "pinia";

interface State {
  baseFolder: string;
}

export const useMainStore = defineStore("main", {
  state: (): State => ({
    baseFolder: "",
  }),
  actions: {
    setBaseFolder(bf: string) {
      this.baseFolder = bf;
    }
  },
});