<script>
import SearchBar from "./SearchBar.vue";
import MenuButton from "./MenuButton.vue";

export default {
	components: {
		SearchBar: SearchBar,
		MenuButton: MenuButton,
	},
}
</script>

<template>
	<div class="container">
		<img src="../assets/images/png/planet.png">
		<SearchBar placeholder="Search and Create" />
		
		<MenuButton text="Select Base Folder" iconType="folder" />
		
		<!-- don't show if there's no folders, or files
			<h3>Recent Folders</h3>
		-->
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
	width: 40%;
	margin-bottom: 10px;
	
	margin-left: 50%;
	transform: translateX(-50%);
}
</style>
