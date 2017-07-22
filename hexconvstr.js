function toHex(d) {
    var r = d % 16;
    var result;
    if (d - r == 0)
        result = toChar(r);
    else
        result = toHex( (d - r)/16 ) + toChar(r);
    return result;
}

function toChar(n) {
    const alpha = "0123456789ABCDEF";
    return alpha.charAt(n);
}