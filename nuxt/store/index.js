import axios from "axios";

export const state = () => ({
    //変数置き
    people: 0,
    monday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    tuesday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    wednesday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    thursday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    friday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    saturday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    sunday:[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
})

export default monday

export const mutations = {

    setPeople(people, num) {
        people = num;
    },
    setGraphMon(monday, vals){
        monday = vals;
    },
    setGraphTue(tuesday, vals){
        tuesday = vals;
    },
    setGraphWed(wednesday, vals){
        wednesday = vals;
    },
    setGraphThu(thursday, vals){
        thursday = vals;
    },
    setGraphFri(friday, vals){
        friday = vals;
    },
    setGraphSat(saturday, vals){
        saturday = vals;
    },
    setGraphSun(sunday, vals){
        sunday = vals;
    },
}

export const actions = {
    async getPeople(context) {
        await axios
            .get('http://localhost:8081/people')
            .then((res) => {
                context.commit("setPeople", res.data);
            })
            .catch(() => {
                console.log("peopelの取得が失敗しました。")
            });
    },
    async getGraphVal(context) {
        await axios
            .get('http://localhost:8081/graph')
            .then((res) => {
                context.commit("setGraphMon", res.data.monday);
                context.commit("setGraphTue", res.data.tuesday);
                context.commit("setGraphWed", res.data.wednesday);
                context.commit("setGraphThu", res.data.thursday);
                context.commit("setGraphFri", res.data.friday);
                context.commit("setGraphSat", res.data.saturday);
                context.commit("setGraphSun", res.data.sunday);
            })
            .catch(() => {
                console.log("graphの取得が失敗しました。")
            });
    },
}

