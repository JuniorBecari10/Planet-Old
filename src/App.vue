<script setup lang="ts">
import { useMainStore } from "./store/main";

const main = useMainStore();
</script>

<script lang="ts">
	import WelcomeScreen from "./components/WelcomeScreen.vue";

	export default {
		components: {
			WelcomeScreen: WelcomeScreen,
		}
	}
</script>

<template>
	<WelcomeScreen />
</template>
