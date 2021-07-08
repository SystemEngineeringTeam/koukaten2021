import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

export const state = () => ({
    //変数置き
    nums: [
        { people: 0, }
    ],
})

export const mutations = {

    setPeole(state, nums) {
        state.nums = nums;
    },
}

export const actions = {
    async getPeople(context) {
        await axios
            .get('http://localhost:8081/people')
            .then((res) => {
                context.commit("setPeople", res.data);//コ↑コ↓　わからん
            })
            .catch(() => {
                console.log("peopelの取得が失敗しました。")
            });
    },
}