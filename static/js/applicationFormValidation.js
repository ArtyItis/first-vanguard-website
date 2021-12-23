function validate(input, bool) {
    if (bool) {
        input.classList.add("is-valid");
        input.classList.remove("is-invalid");
        return true;
    } else {
        input.classList.add("is-invalid");
        input.classList.remove("is-valid");
        return false;
    }
}
function isCheckboxGroupValid(wrapper_element_id, invalid_id) {
    let inputs = document.getElementById(wrapper_element_id).getElementsByTagName("input");
    let invalid = document.getElementById(invalid_id);
    for (let input of inputs) {
        if (input.checked) {
            invalid.classList.add("invisible");
            return true;
        }
    }
    invalid.classList.remove("invisible");
    return false;
}
//------------ Guidelines
function isCompanyGuidelineValid(inputId) {
    let input = document.getElementById(inputId);
    let invalid = document.getElementById("company-guidelines-invalid");
    if (input.checked == false) {
        invalid.classList.remove("invisible");
        return false;
    } else {
        invalid.classList.add("invisible");
        return true;
    }
}
//------------General
function validateCharacterName() {
    let input = document.getElementsByName("charactername")[0];
    let bool = input.value != "";
    return validate(input, bool);
}

function validateDiscordTag() {
    let input = document.getElementsByName("discord-tag")[0];
    let regex = new RegExp("^.+#.+$");
    let bool = regex.test(input.value);
    return validate(input, bool);
}

function validateGearScore() {
    let input = document.getElementsByName("gearscore")[0];
    let regex = new RegExp("^[0-9]*$");
    let bool = regex.test(input.value);
    if (bool) {
        bool = 0 < input.value && input.value <= 600;
    }
    return validate(input, bool);
}

function validateExpertise() {
    let input = document.getElementsByName("expertise")[0];
    let regex = new RegExp("^[0-9]*$");
    let bool = regex.test(input.value);
    if (bool) {
        bool = 0 < input.value && input.value <= 590;
    }
    return validate(input, bool);
}

function validateArmorWeight() {
    let input = document.getElementsByName("armor-weight")[0];
    let bool = input.value != "";
    return validate(input, bool);
}

//------------Attributes
function validateAttributes(){
    let inputArray = [5];
    inputArray[0] = document.getElementsByName("strength")[0];
    inputArray[1] = document.getElementsByName("dexterity")[0];
    inputArray[2] = document.getElementsByName("intelligence")[0];
    inputArray[3] = document.getElementsByName("focus")[0];
    inputArray[4] = document.getElementsByName("constitution")[0];
    let regex = new RegExp("^[0-9]*$");
    let boolArray = [5];
    for (let index = 0; index < 5; index++) {
        boolArray[index] = regex.test(inputArray[index].value) && inputArray[index].value > 4;
    }
    let sumPoints = 0;
    for (let i = 0; i < 5; i++) {
        sumPoints = Number(inputArray[i].value) + sumPoints;
    }
    if (sumPoints < 25 || 475 < sumPoints) {
        for (let j = 0; j < 5; j++) {
            boolArray[j] = false;
        }
    }
    console.log(sumPoints);
    validate(inputArray[0], boolArray[0]);
    validate(inputArray[1], boolArray[1]);
    validate(inputArray[2], boolArray[2]);
    validate(inputArray[3], boolArray[3]);
    validate(inputArray[4], boolArray[4]);
    let sumBool = true;
    for (let index = 0; index < 5; index++) {
        sumBool = boolArray[index] && sumBool;
    }
    return sumBool;
}

//------------Form
function validateForm() {
    // let company_taxes = isCompanyGuidelineValid("company-taxes");
    let pvp_activity = isCompanyGuidelineValid("pvp-activity");
    let equipment_flexibility = isCompanyGuidelineValid("equipment-flexibility");
    let discord_activity = isCompanyGuidelineValid("discord-activity");
    let wars = isCompanyGuidelineValid("wars");
    let progress = isCompanyGuidelineValid("progress");

    let company_guidelines = pvp_activity && equipment_flexibility && discord_activity && wars && progress;

    let name = validateCharacterName();
    let dc = validateDiscordTag();
    let gs = validateGearScore();
    let expertise = validateExpertise();
    let armorWeight = validateArmorWeight();

    let attributes = validateAttributes();

    let roles = isCheckboxGroupValid("roles", "roles-invalid");
    let weapons = isCheckboxGroupValid("weapons", "weapons-invalid");

    let isValid = company_guidelines && name && dc && gs && expertise && armorWeight && attributes && roles && weapons;
    let form = document.getElementById("application-form");
    if (isValid) {
        form.submit();
    }
}