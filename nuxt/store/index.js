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
        async getPeple(people){
            await axios
            .get('http://localhost:')
        },
    },
});