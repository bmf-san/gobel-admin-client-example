<template>
  <div class="container-readable">
    <h1>Home</h1>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "Home"
});
</script>

<style>
.multiselect__tags {
  height: 5rem;
  border-radius: 0.3rem;
  background-color: #fff;
  border: 0.1rem solid var(--border-color);
  width: 100%;
  font-size: 1.6rem;
  padding: 1rem;
  margin-bottom: 2.2rem;
  -webkit-appearance: none;
  appearance: none;
  -moz-appearance: none;
}

.multiselect__input {
  font-size: 1.4rem;
  padding: 0;
  border: none;
  border-radius: 0;
  height: 2.6rem;
  outline: none;
  margin-top: 2.2rem;
  margin-right: 0rem;
  margin-bottom: 0rem;
  margin-left: 0rem;
  caret-color: transparent;
}

.multiselect__content {
  padding: 0;
  margin-bottom: 2.2rem;
}
.multiselect__input:focus {
  background-color: transparent;
}

.multiselect__content {
  list-style-type: none;
}

.multiselect__element,
.multiselect__tag {
  padding: 0.5rem 1.4rem;
  line-height: 1;
  display: inline-block;
  background: #fff;
  text-decoration: none;
  color: var(--main-color);
  border: var(--main-color) solid 0.1rem;
  font-size: 1.4rem;
  height: 2.6rem;
  margin-right: 1rem; /* additonal style */
}
.multiselect__tag-icon:after {
  content: "\D7";
  color: #266d4d;
  font-size: 14px;
}
</style>
