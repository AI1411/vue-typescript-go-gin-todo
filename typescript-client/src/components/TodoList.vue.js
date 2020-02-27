import { __decorate } from "tslib";
import { Component, Vue } from 'vue-property-decorator';
import axios from 'axios';
const NOT_STARTED = 1;
const FINISHED = 3;
let TodoList = class TodoList extends Vue {
    constructor() {
        super(...arguments);
        this.todoList = [];
        this.inputField = '';
        this.baseUrl = 'http://localhost:8080/api/v1/';
    }
    created() {
        this.getTodo();
    }
    async getTodo() {
        try {
            const response = await axios.get(this.baseUrl + 'todo');
            this.todoList = response.data;
            return this.todoList;
        }
        catch (e) {
            return e;
        }
    }
    async addTodo() {
        if (!this.inputField) {
            return;
        }
        try {
            const params = {
                text: this.inputField,
                status: 1,
            };
            await axios.post(this.baseUrl + 'todo', JSON.stringify(params));
            this.getTodo();
            this.inputField = '';
        }
        catch (e) {
            return e;
        }
    }
    async deleteTodo(todo) {
        try {
            await axios.delete(this.baseUrl + 'todo/' + todo.ID);
            this.getTodo();
        }
        catch (e) {
            return e;
        }
    }
    async toggle(todo) {
        try {
            let status = 0;
            if (todo.Status === NOT_STARTED) {
                status = FINISHED;
            }
            else {
                status = NOT_STARTED;
            }
            const params = {
                '{status}': status,
            };
            await axios.put(this.baseUrl + 'todo/' + todo.ID, JSON.stringify(params));
            todo.Status = status;
        }
        catch (e) {
            return e;
        }
    }
};
TodoList = __decorate([
    Component
], TodoList);
export default TodoList;
//# sourceMappingURL=TodoList.vue.js.map