<script setup lang="ts">
	import { ref, Ref } from "vue";

	var baseFolder: Ref<string> = ref("");
	
	function setBaseFolder(path: string): void {

	}

	var functions: Ref<{}> = ref({
		setBaseFolder: setBaseFolder,
	});
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
