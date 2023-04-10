<script setup>
defineProps({
  msg: {
    type: String,
  }
})
</script>

<template>
  <div class="status">
    <h1 class="orange">Putte: {{ msg }}</h1>
    <h3>
      This would be the main status page!
    </h3>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.status h1,
.status h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .status h1,
  .status h3 {
    text-align: left;
  }
}
</style>
