import re


def reverse_words(sentence: str) -> str:
    sentence = re.sub(" +", ' ', sentence.strip())

    sentence =  list(sentence)
    str_len = len(sentence) - 1

    str_rev(sentence, 0, str_len)
    start = 0

    for end in range(0, str_len + 1):
        if end == str_len or sentence[end] == ' ':
            end_idx = end if end == str_len else end - 1
            str_rev(sentence, start, end_idx)
            start = end + 1

    return ''.join(sentence)





def str_rev(_str, start, end):
    while start < end:
        _str[start], _str[end] = _str[end], _str[start]
        start += 1
        end -= 1

