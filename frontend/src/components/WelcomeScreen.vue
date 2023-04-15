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

		<div class="button">
			<a href="">
				<img src="../assets/images/svg/new-folder.svg">
				<span>Open Base Folder</span>
			</a>
		</div>

		<h3>Recent Folders</h3>
	</div>
</template>

<style scoped>
	div.container {
		width: 30%;

		margin-left: 50%;
		transform: translateX(-50%);

		display: flex;

		align-items: left;
		justify-content: center;
		flex-direction: column;
	}

	div.container > *:not(img) {
		margin-top: 20px;
	}

	.container > img {
		width: 50%;

		margin-left: 50%;
		transform: translateX(-50%);
	}

	.button {
		width: 100%;
	}

	.button a {
		display: flex;
		width: 100%;

		align-items: center;
	}

	.button img {
		width: 32px;
		margin-right: 20px;
	}
</style>
