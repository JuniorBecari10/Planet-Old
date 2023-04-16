<script setup>
import { OpenDirDialog } from "../../wailsjs/go/main/App";

function openDialog() {
	OpenDirDialog('Select Base Folder', true)
	.then(res => {
		if (res === "") return;

		alert(res);
	});
}
</script>

<template>
  <div class="button ripple" @click="() => { openDialog(); }">
			<a href="">
				<img src="../assets/images/svg/new-folder.svg">
				<span>Select Base Folder</span>
			</a>
		</div>
</template>

<style>
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