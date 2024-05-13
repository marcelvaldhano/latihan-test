function isPalindrom(str){
    let newStr=str.toLowerCase().replace(/[\W_]/g,'')
    let reverseStr=newStr.split('').reverse().join('')
    return newStr===reverseStr
}