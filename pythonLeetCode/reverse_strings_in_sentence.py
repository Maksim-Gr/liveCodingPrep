import re


def reverse_words(sentence: str) -> str:
    sentence = re.sub(" +", ' ', sentence.strip())

    sentence =  list(sentence)
    str_len = len(sentence) - 1

