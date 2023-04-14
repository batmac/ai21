package main

import (
	"fmt"

	"github.com/batmac/ai21"
)

func main() {
	client := ai21.NewClientFromEnv()
	/* resp, err := client.ContextualAnswers(ContextualAnswersRequest{
		Question: "Did the economy shrink after the Omicron variant arrived?",
		Context:  "In 2020 and 2021, enormous QE — approximately $4.4 trillion, or 18%, of 2021 gross domestic product (GDP) — and enormous fiscal stimulus (which has been and always will be inflationary) — approximately $5 trillion, or 21%, of 2021 GDP — stabilized markets and allowed companies to raise enormous amounts of capital. In addition, this infusion of capital saved many small businesses and put more than $2.5 trillion in the hands of consumers and almost $1 trillion into state and local coffers. These actions led to a rapid decline in unemployment, dropping from 15% to under 4% in 20 months — the magnitude and speed of which were both unprecedented. Additionally, the economy grew 7% in 2021 despite the arrival of the Delta and Omicron variants and the global supply chain shortages, which were largely fueled by the dramatic upswing in consumer spending and the shift in that spend from services to goods. Fortunately, during these two years, vaccines for COVID-19 were also rapidly developed and distributed.In today's economy, the consumer is in excellent financial shape (on average), with leverage among the lowest on record, excellent mortgage underwriting (even though we've had home price appreciation), plentiful jobs with wage increases and more than $2 trillion in excess savings, mostly due to government stimulus. Most consumers and companies (and states) are still flush with the money generated in 2020 and 2021, with consumer spending over the last several months 12% above pre-COVID-19 levels. (But we must recognize that the account balances in lower-income households, smaller to begin with, are going down faster and that income for those households is not keeping pace with rising inflation.) Today's economic landscape is completely different from the 2008 financial crisis when the consumer was extraordinarily overleveraged, as was the financial system as a whole — from banks and investment banks to shadow banks, hedge funds, private equity, Fannie Mae and many other entities. In addition, home price appreciation, fed by bad underwriting and leverage in the mortgage system, led to excessive speculation, which was missed by virtually everyone — eventually leading to nearly $1 trillion in actual losses.",
	}) */ /*
		resp, err := client.Summarize(SummarizeRequest{
			Source:     "We’ve all experienced reading long, tedious, and boring pieces of text - financial reports, legal documents, or terms and conditions (though, who actually reads those terms and conditions to be honest?).  Imagine a company that employs hundreds of thousands of employees. In today's information overload age, nearly 30% of the workday is spent dealing with documents. There's no surprise here, given that some of these documents are long and convoluted on purpose (did you know that reading through all your privacy policies would take almost a quarter of a year?). Aside from inefficiency, workers may simply refrain from reading some documents (for example, Only 16% of Employees Read Their Employment Contracts Entirely Before Signing!).   This is where AI-driven summarization tools can be helpful: instead of reading entire documents, which is tedious and time-consuming, users can (ideally) quickly extract relevant information from a text. With large language models, the development of those tools is easier than ever, and you can offer your users a summary that is specifically tailored to their preferences.  Let's take legal documents, for example. Though they are written in English, many people find legal documents to be difficult to comprehend, as if they were actually written in a foreign language. Moreover, the interesting parts of each document may differ depending on the person who reads it, so off-the-shelf summarization tools may be too general or too specific. As an example, let's look at the involved personas:",
			SourceType: SourceTypeText,
		}) */
	/*
		resp, err := client.Complete(ModelGrande, CompleteRequest{
			Prompt: "what is a capybara",
			/* 		NumResults:       0,
			   		MaxTokens:        0,
			   		MinTokens:        0,
			   		Temperature:      0,
			   		TopP:             0,
			   		StopSequences:    []string{},
			   		TopKReturn:       0,
			   		FrequencyPenalty: Penalty{},
			   		PresencePenalty:  Penalty{},
			   		CountPenalty:     Penalty{},
		}) */

	resp, err := client.EasyComplete(ai21.ModelJumbo, "what is a capybara", 50)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", resp)
}
