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

function isCompanyGuidelineValid(input_id, collapse_id) {
    var input = document.getElementById(input_id);
    var collapse = document.getElementById(collapse_id);
    var invalid = document.getElementById("company-guidelines-invalid");
    if (input.checked == false) {
        invalid.classList.remove("invisible");
        return false;
    } else {
        invalid.classList.add("invisible");
        return true;
    }
}

function isInputTextValid(input_id, invalid_id) {
    var input = document.getElementsByName(input_id)[0];
    var invalid = document.getElementById(invalid_id);
    if (input.value == "") {
        invalid.classList.remove("invisible");
        return false;
    } else {
        invalid.classList.add("invisible");
        return true;
    }
}

function validateForm() {
    //company guidelines
    var company_taxes = isCompanyGuidelineValid("company-taxes", "collapseOne");
    var pvp_activity = isCompanyGuidelineValid("pvp-activity", "collapseTwo");
    var equipment_flexibility = isCompanyGuidelineValid("equipment-flexibility", "collapseThree");
    var discord_activity = isCompanyGuidelineValid("discord-activity", "collapseFour");
    //character
    var charactername = isInputTextValid("charactername", "character-invalid");
    var gearscore = isInputTextValid("gearscore", "character-invalid");
    //attributes
    var strength = isInputTextValid("strength", "attributes-invalid");
    var dexterity = isInputTextValid("dexterity", "attributes-invalid");
    var intelligence = isInputTextValid("intelligence", "attributes-invalid");
    var focus = isInputTextValid("focus", "attributes-invalid");
    var constitution = isInputTextValid("constitution", "attributes-invalid");
    //summary
    var company_guidelines = company_taxes && pvp_activity && equipment_flexibility && discord_activity;
    var attributes = strength && dexterity && intelligence && focus && constitution;
    //roles
    var roles = isCheckboxGroupValid("roles", "roles-invalid");
    //weapons
    var weapons = isCheckboxGroupValid("weapons", "weapons-invalid");

    var form = document.getElementById("recruitment-form");
    if (company_guidelines && charactername && gearscore && attributes && weapons && roles) {
        form.submit();
    }
}