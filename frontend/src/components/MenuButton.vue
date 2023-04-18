<script setup>
import { OpenDirDialog } from "../../wailsjs/go/main/App";
import { global } from "../main.js";

// all the functions a button can have

function selectBaseFolder() {
	OpenDirDialog()
	.then(res => {
		if (res === "") return;
		global.baseFolder = res;

		alert(global.baseFolder);
	});
}
</script>

<script>
  export default {
    props: {
      text: String,
      iconType: String,
    }
  }
</script>

<template>
  <div class="button ripple" @click="() => selectBaseFolder()">
			<a href="">
				<img v-if="iconType === 'folder'" src="../assets/images/svg/folder.svg">
				<span>{{ text }}</span>
				{{ global.baseFolder }}
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