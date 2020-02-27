<template>
    <div class="home">
        <img src="../assets/logo.png" alt="">
        <TodoList/>
    </div>
</template>

<script lang="ts">
    import {Component, Vue} from 'vue-property-decorator'
    import TodoList from '@component/TodoList.vue';

    @Component({
        components: {
            TodoList
        },
    })
    export default class Todo extends Vue {

    }
</script>

<style scoped>

</style>