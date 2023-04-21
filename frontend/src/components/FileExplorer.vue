<script setup>
import { mapActions, storeToRefs } from 'pinia';
import { useMainStore } from '../store/main';

const main = useMainStore();

const { baseFolder } = storeToRefs(main);
const { setBaseFolder } = mapActions(useMainStore, ["increment"]);
</script>

<template>
  aaaaaaaaaaaa
  {{ baseFolder }}
</template>