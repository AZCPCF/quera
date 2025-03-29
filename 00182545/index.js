function encrypt(str, n) {
    return str
        .split('')
        .map(char => shiftChar(char, n))
        .join('');
}

function decrypt(str, n) {
    return str
        .split('')
        .map(char => shiftChar(char, -n))
        .join('');
}

function shiftChar(char, shift) {
    const isUpperCase = char >= 'A' && char <= 'Z';
    const isLowerCase = char >= 'a' && char <= 'z';

    if (isUpperCase || isLowerCase) {
        const baseCode = isUpperCase ? 'A'.charCodeAt(0) : 'a'.charCodeAt(0);
        const charCode = char.charCodeAt(0) - baseCode;
        const newCode = (charCode + shift + 26) % 26; 
        return String.fromCharCode(newCode + baseCode);
    }

    return char; 
}

export { encrypt, decrypt };
