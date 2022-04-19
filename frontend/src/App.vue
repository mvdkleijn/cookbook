<template>
  <link href="https://unpkg.com/primevue/resources/themes/saga-blue/theme.css " rel="stylesheet">
  <link href="https://unpkg.com/primevue/resources/primevue.min.css " rel="stylesheet">
  <link href="https://unpkg.com/primeicons/primeicons.css " rel="stylesheet">
  <div>
    <router-view />
  </div>
</template>

<style lang="scss">
* {
  margin: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
html, body {
  margin: 0;
  height: 100%;
}

.container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.header {
  flex-shrink: 0;
  height: 55px;
}

.content {
  flex: 1;
  overflow-y: scroll;
}
</style>
