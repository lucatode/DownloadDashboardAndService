<template>
  <!--How to include the question here???-->
  <input :placeholder="placeholder">
</template>
<script>
  import SurveyInputBase from './SurveyInputBase.vue';
  export default {
    extends: SurveyInputBase,
    props: [ 'placeholder' ],
  }
</script>