<script>
  export default {
    props: {
      placeholder: String,
    }
  }
</script>

<template>
  <div>
    <img src="../assets/images/svg/search-icon.svg">
    <input type="text" :placeholder="placeholder">
    <span>{{ placeholder }}</span>
  </div>
</template>

<style scoped>
  div {
    background: #22252C;
    padding: 8px;

    display: flex;
    width: 100%;

    border-radius: 8px;
  }

  img {
    margin-right: 20px;
  }

  input {
    background: transparent;

    border: none;
    outline: none;

    width: 100%;
    color: #fff;
  }
</style>