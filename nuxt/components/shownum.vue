<!--人数表示コンポーネント-->
<template>
    <div>
        <!--タイトル-->
        <h3>現在の人数</h3>
        <div>
            <!--人数表示部分-->
            <h3>現在：</h3><h2>{{this.$store.state.tasks}}</h2><h3>人</h3>
        </div>
    </div>
</template>

<script>
export default {
    name: "ShowNum"
}
</script>

<style scoped>
    h2{
        display: inline;
    }
    h3{
        display: inline;
    }
</style>