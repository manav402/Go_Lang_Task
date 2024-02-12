
// the frequently used variable like input box and error box are declared as global

const inputBox = document.querySelector("input");
const errorBox = document.querySelector(".error");
let isCalculated = false;

// to maintain code readebility created const variable with asssociate integer

const EQUAL = 0;
const BACKSPACE = 1;
const CLEAR = 2;

// a global warning function called from anywhere to show a red alert to user on any type of error
// @params :- str :- a string containing the warning massage

function warn(str){
	errorBox.innerText =str;
	inputBox.value=null;
	document.querySelector('button[id="clear"]').focus();
}


// js float precision is not up to mark making the funciton return a round off value based on input
// @params :- floatstr :- answer string of eval function which has float pointer and length is higher than 5
// @return :- val :- returns the answer precised by the function 

function floatPrecision(floatStr){
	try {
		let floatNumber = parseFloat(floatStr).toPrecision(5);
		return floatNumber;
	} catch (e) {
		warn("can't parse the float string error:- " + e);
		return floatStr;
	}
}

// evaluate function which evaluate the equation from calculator input box

function evaluate() {

	// testing the input with the regex to authenticate the user iputs

	if(!isCalculated){			
		let regexSolver = /([\+\-]?\d+\.?\d*[\+\-\*\/]?)/.test(inputBox.value);

		if (regexSolver) {
			let val = eval(inputBox.value)
			
			val = (val - parseInt(val) != 0) ? floatPrecision(val) : val;
			
			updateInput(val, EQUAL);			
		} else {
			warn("please enter a valid equation \n press enter to try again");
		}

		isCalculated = true;
	}else{
		isCalculated = false;
		updateInput(null,CLEAR);
	}
}

// updater function connected with input box reflects the users input
// @params:- input - any value or sign will be considered as input
//  signal - any other key such as delete or equal button will be passed as signal
function updateInput(input, signal) {

	let value = inputBox.value;
	// performing necessary actions according to the signal value
	switch (signal) {
		case BACKSPACE:
			inputBox.value = value.slice(0, value.length - 1);
			break;
		case EQUAL:
			inputBox.value = input;
			break;
		case CLEAR:
			isCalculated = false;
			inputBox.value = null;
			errorBox.innerText = null;
			break;
		default:
			inputBox.value += input;
			break;
	}

}

// listener function listen to user inputs and binded with each buttons event listeners
//  this binding represent the buttons of html
function listener() {
	let value = this.value;
	// taking action based on the user clicked button value
	if(isCalculated){
		updateInput(null,CLEAR)
	}
	switch (value) {
		case "equal":
			try{
				evaluate();
			}catch(e){
				warn("error :- can't parse your string please try again \n please press enter " + e);
			}
			break;
		case "backSpace":
			updateInput(null, BACKSPACE);
			break;
		case "clear":
			updateInput(null, CLEAR);
			break;
		default:
			updateInput(value, null);
			break;
	}
}

// looping through array of button and binding event listener to each of them representing a clean html ui.
let btnArray = document.querySelectorAll("button");
btnArray.forEach((element) => {
	element.addEventListener("click", listener);
});
