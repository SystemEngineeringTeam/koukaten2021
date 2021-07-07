<!--人数表示コンポーネント-->
<template>
    <div>
        <!--タイトル-->
        <h3>現在の人数</h3>
        <div>
            <!--人数表示部分-->
            <h2>
                {{this.$store.state.tasks}}
            </h2>
        </div>
    </div>
</template>
<script>
export default {
    
}
</script>