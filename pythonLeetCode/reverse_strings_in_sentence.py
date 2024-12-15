import re


def reverse_words(sentence: str) -> str:
    sentence = re.sub(" +", ' ', sentence.strip())

    sentence =  list(sentence)
    str_len = len(sentence) - 1





def str_rev(_str, start, end):
    while start < end:
        if _str[start] != ' ':
            break
        start += 1