import axios from 'axios';

class Requests {

    getCalculating() {
        return axios.get('http://localhost:8080/api/v1/stop');
    }
    async getState(tilesNum, he, solvable) {
        return axios.get(`http://localhost:8080/api/v1/state/${tilesNum}/${he}/${solvable}`);
    }

    // startAlgo() {
    //     return axios.get('http://localhost:8080/api/v1/algo');
    // }

    getPath() {
        return axios.get('http://localhost:8080/api/v1/path');
    }

    async putState(State) {
        console.log('put firstState');
        console.log(State);
        return axios.put('http://localhost:8080/api/v1/state', State);
    }


}

export default new Requests();