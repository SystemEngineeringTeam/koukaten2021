import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
    //変数置き
    state:{
        tasks:[
            {people: 0,}
        ],
    },

    mutations:{
        setPeole(state, tasks){
            state.tasks = tasks;
        },
    },

    actions: {
        async getPeople(people){
            await axios
            .get('http://localhost:8081/people')
            .then((res) => {
                context.commit("setPeople", );
            })
            .catch(() => {
                console.log("peopelの取得が失敗しました。")
            });
        },
    },
});