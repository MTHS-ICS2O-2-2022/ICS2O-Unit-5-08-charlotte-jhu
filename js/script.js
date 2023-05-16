// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function myButtonClicked() {
  // This function does division using a loop
  // input
  let number1 = parseInt(document.getElementById("number-1").value)
  const number2 = parseInt(document.getElementById("number-2").value)

  // process
  let counter = 0
  for (counter <= number2; counter ++;) {
    number1 -= number1
  }

  // output
  document.getElementById("answer").innerHTML =
    number1 + " divided by " + number2 + " is " + counter

  // remainder output
  if (number1 != 0)
  document.getElementById("remainder").innerHTML = "With a remainder of " + number1
}
