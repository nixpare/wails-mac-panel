import {GreetService} from "./bindings/gui";
import {Events} from "@wailsio/runtime";

const resultElement = document.getElementById('result');
const timeElement = document.getElementById('time');

window.doGreet = async () => {
    let name = document.getElementById('name').value;
    if (!name || name === "") {
        name = 'Anonymous';
    }
    
    const result = await GreetService.Greet(name)
        .catch((err) => {
            console.log(err);
        });
    if (!result) return;

    resultElement.innerText = result;
}

Events.On('time', (time) => {
    timeElement.innerText = time.data;
});
