import axios from 'axios'

export const state = () => ({
  //変数置き
  people: 0,
  monday: [
    10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 3, 4, 6, 3, 3, 2, 3, 4, 2, 2, 1, 1, 0, 0,
  ],
  tuesday: [
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 2, 0, 0, 0, 0, 0, 0,
  ],
  wednesday: [
    0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 2, 1, 2, 0, 0, 0, 0, 0, 0,
  ],
  thursday: [
    0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 3, 3, 2, 1, 1, 1, 3, 3, 2, 0, 0, 0, 0,
  ],
  friday: [
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 4, 3, 3, 1, 0, 0, 0, 0, 0, 0, 0, 0,
  ],
  saturday: [
    0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
  ],
  sunday: [
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
  ],
  date: '2021-10-01 10:00:00',
})

export const mutations = {
  setPeople(state, num) {
    state.people = num
  },
  setGraphMon(state, vals) {
    state.monday = vals
  },
  setGraphTue(state, vals) {
    state.tuesday = vals
  },
  setGraphWed(state, vals) {
    state.wednesday = vals
  },
  setGraphThu(state, vals) {
    state.thursday = vals
  },
  setGraphFri(state, vals) {
    state.friday = vals
  },
  setGraphSat(state, vals) {
    state.saturday = vals
  },
  setGraphSun(state, vals) {
    state.sunday = vals
  },
  setDate(state, date) {
    state.date = date
  },
}

export const actions = {
  async getPeople(context) {
    await axios
      .get('http://localhost:8081/people')
      .then((res) => {
        context.commit('setPeople', res.data.people) //コ↑コ↓　わからん
        context.commit('setDate', res.data.date)
      })
      .catch(() => {
        console.log('peopelの取得に失敗しました。')
      })
  },
  async getGraphVal(context) {
    await axios
      .get('http://localhost:8081/graph')
      .then((res) => {
        context.commit('setGraphMon', res.data.monday)
        context.commit('setGraphTue', res.data.tuesday)
        context.commit('setGraphWed', res.data.wednesday)
        context.commit('setGraphThu', res.data.thursday)
        context.commit('setGraphFri', res.data.friday)
        context.commit('setGraphSat', res.data.saturday)
        context.commit('setGraphSun', res.data.sunday)
      })
      .catch(() => {
        console.log('graphの取得に失敗しました。')
      })
  },
}
