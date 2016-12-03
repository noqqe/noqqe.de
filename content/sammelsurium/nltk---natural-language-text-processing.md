---
title: NLTK - Natural Language Text Processing
date: 2016-04-13T17:01:08
tags: 
- python
- Programming
- nltk
- text
- analysis
---

Sentiment Analysis - Ist der Text Positiv oder negativ?

Words per Sentence

    len(nltk.corpus.brown.words()) / len(nltk.corpus.brown.sents())

Anzahl der Wörter finden

		len(nltk.tokenize.sent_tokenize(foo))

Anzahl der Sätze finden

		len(nltk.tokenize.sent_tokenize(text1))

Search text for every occurence of word:

    text1.concordance('foo')

Find synonyms:

    text1.similar('foo')

Generate random text:

    text1.generate()

Count Vocabulary:

    len(text1)

Count occurance of word:

    text1.count('foo')

See all words from text:

    set(text1)

Hapaxes (least common words):

    FreqDist(text1).hepaxes()

Find Long Words:

    sorted([w for w in set(text1) in len(w) > 15])

#### Links

http://nltk-trainer.readthedocs.org/en/latest/train_classifier.html
