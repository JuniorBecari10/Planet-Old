import { defineStore } from "pinia";

export const useMainStore = defineStore("main", {
  state: () => ({
    baseFolder: "",
  }),
  actions: {
    setBaseFolder(bf: string) {
      this.baseFolder = bf;
    }
  },
});