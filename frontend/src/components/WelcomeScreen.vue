<script setup>
	
</script>

<script>
	import SearchBar from "./SearchBar.vue";

	export default {
		components: {
			SearchBar: SearchBar,
		}
	}
</script>

<template>
	<div class="container">
		<img src="../assets/images/png/planet.png">
		<SearchBar placeholder="Search and Create" />
	</div>
</template>

<style scoped>
	div.container {
		width: 30%;

		margin-left: 50%;
		transform: translateX(-50%);

		display: flex;

		align-items: center;
		justify-content: center;
		flex-direction: column;
	}

	img {
		width: 50%;
		margin-bottom: 20px;
	}
</style>
