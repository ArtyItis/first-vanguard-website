function isCheckboxGroupValid(wrapper_element_id, invalid_id) {
    var inputs = document.getElementById(wrapper_element_id).getElementsByTagName("input");
    var invalid = document.getElementById(invalid_id);
    for (let input of inputs) {
        if (input.checked) {
            invalid.classList.add("invisible");
            return true;
        }
    }
    invalid.classList.remove("invisible");
    return false;
}

function isCompanyGuidelineValid(inputId) {
    var input = document.getElementById(inputId);
    var invalid = document.getElementById("company-guidelines-invalid");
    if (input.checked == false) {
        invalid.classList.remove("invisible");
        return false;
    } else {
        invalid.classList.add("invisible");
        return true;
    }
}

function isInputTextValid(inputName, invalidId) {
    var input = document.getElementsByName(inputName)[0];
    var invalid = document.getElementById(invalidId);
    if (input.value == "") {
        invalid.classList.remove("invisible");
        return false;
    } else {
        invalid.classList.add("invisible");
        return true;
    }
}

function containsNumber(inputName) {
    var input = document.getElementsByName(inputName)[0];
    var regex = new RegExp("^[0-9]*$");
    if (input.value != "" && regex.test(input.value)) {
        return true;
    } else {
        return false;
    }
}

function validateGeneralInformation() {
    var charactername = document.getElementsByName("charactername")[0];
    var discordTag = document.getElementsByName("discord-tag")[0];
    var armorWeight = document.getElementsByName("armor-weight")[0];
    var vC = charactername.value != "";
    console.log(document.getElementsByName("gearscore")[0].value);
    var vG = containsNumber("gearscore") && (document.getElementsByName("gearscore")[0].value >= 0 && document.getElementsByName("gearscore")[0].value <= 600);
    var vW = armorWeight.value != "";
    var vD = discordTag.value.includes("#");

    var invalid = document.getElementById("character-invalid");

    if (vC && vG && vW && vD) {
        invalid.classList.add("invisible");
        return true;
    } else {
        invalid.classList.remove("invisible");
        return false;
    }
}

function validateAttributes() {
    var vS = containsNumber("strength")
    var vD = containsNumber("dexterity");
    var vI = containsNumber("intelligence");
    var vF = containsNumber("focus");
    var vC = containsNumber("constitution");
    var invalid = document.getElementById("attributes-invalid");
    var result = false;
    if (vS && vD && vI && vF && vC) {
        var strength = document.getElementsByName("strength")[0].value;
        var dexterity = document.getElementsByName("dexterity")[0].value;
        var intelligence = document.getElementsByName("intelligence")[0].value;
        var focus = document.getElementsByName("focus")[0].value;
        var constitution = document.getElementsByName("constitution")[0].value;
        var sum = strength + dexterity + intelligence + focus + constitution;
        if (25 <= sum && sum <= 475) {
            result = true;
        }
    }
    if (result) {
        invalid.classList.add("invisible");
        return true;
    } else {
        invalid.classList.remove("invisible");
        return false;
    }
}

function validateForm() {
    //company guidelines
    var company_taxes = isCompanyGuidelineValid("company-taxes");
    var pvp_activity = isCompanyGuidelineValid("pvp-activity");
    var equipment_flexibility = isCompanyGuidelineValid("equipment-flexibility");
    var discord_activity = isCompanyGuidelineValid("discord-activity");

    var company_guidelines = company_taxes && pvp_activity && equipment_flexibility && discord_activity;

    var general = validateGeneralInformation();
    var attributes = validateAttributes();
    var roles = isCheckboxGroupValid("roles", "roles-invalid");
    var weapons = isCheckboxGroupValid("weapons", "weapons-invalid");

    var form = document.getElementById("application-form");
    if (company_guidelines && general && attributes && weapons && roles) {
        form.submit();
    }
}