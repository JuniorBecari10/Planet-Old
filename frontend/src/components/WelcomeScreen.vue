<script setup>
function openDialog() {
	OpenDirDialog('Select Base Folder', true)
	.then(res => {
		if (res === "") return;

		alert(res);
	});
}
</script>

<script>
import SearchBar from "./SearchBar.vue";
import { OpenDirDialog } from "../../wailsjs/go/main/App";

export default {
	components: {
		SearchBar: SearchBar,
	},
}
</script>

<template>
	<div class="container">
		<img src="../assets/images/png/planet.png">
		<SearchBar placeholder="Search and Create" />
		
		<div class="button ripple" @click="() => { openDialog(); }">
			<a href="">
				<img src="../assets/images/svg/new-folder.svg">
				<span>Select Base Folder</span>
			</a>
		</div>
		
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

.button {
	width: 100%;
	padding: 10px;
	border-radius: 10px;
	
	transition: background 0.3s;
}

.button:hover {
	background: #22252C;
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
