package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Theme string
	Points int
	MinimumPoints int
	TimeToResolution int
	Questions []Question
}

func (g *GameState) Init(ch chan bool) {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Println("Escreva o seu nome: ")
	readerName := bufio.NewReader(os.Stdin)

	name, err := readerName.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	fmt.Print("Escolha o tema desse quiz:\n")
	fmt.Print("1 - Geografia\n")
	fmt.Print("2 - História\n")
	fmt.Print("3 - Matemática\n")
	readerTheme := bufio.NewReader(os.Stdin)
	theme, err := readerTheme.ReadString('\n')
	
	
	parsedTheme, _ := toInt(theme[:len(theme)-1])

	switch parsedTheme {
		case 1:
			g.Theme = "geografia"
		case 2:
			g.Theme = "historia"
		case 3:
			g.Theme = "matematica"
		default:
			panic("Opção invalida, selecione outro tema")
	}

	fmt.Printf("Vamos ao jogo %s\n", g.Name)
	ch <- true
}

func (g *GameState) ProcessCSV(ch, csvLoadCh chan bool) {
	<-ch
	csvFile := g.Theme + ".csv"
	fmt.Println(csvFile)
	f, err := os.Open(csvFile)

	if err != nil {
		panic("Erro ao ler arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler csv")
	}

	minimumPoints, _ := toInt(records[1][6])
	timeToResolution, _ := toInt(records[1][7])
	g.MinimumPoints = minimumPoints
	g.TimeToResolution = timeToResolution
	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: correctAnswer,
			}
			g.Questions = append(g.Questions, question)
		}
	}
	csvLoadCh <- true
}

func (g *GameState) Run(csvLoadCh chan bool) {
	<-csvLoadCh
	timeout := time.After(time.Duration(g.TimeToResolution) * time.Second)

gameLoop:
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d %s\033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s \n", j+1, option)
		}

		fmt.Println("Digite uma alternativa: ")

		inputCh := make(chan int)
		go func() {
			for {
				reader := bufio.NewReader(os.Stdin)
				read, _ := reader.ReadString('\n')

				answer, err := toInt(read[:len(read)-1])
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				inputCh <- answer
				return
			}
		}()

		select {
		case answer := <-inputCh:
			if answer == question.Answer {
				fmt.Println("Parabéns você acertou!!!")
				g.Points += 10
			} else {
				fmt.Println("Ops! Errou!")
				fmt.Println("----------------------")
			}
		case <-timeout:
			fmt.Println("\n⏰ Tempo esgotado! O quiz foi encerrado.")
			break gameLoop
		}
	}

	g.CheckScore()
}

func (g* GameState) CheckScore() {
	if g.Points > g.MinimumPoints {
		fmt.Printf("Você foi aprovado com a pontuacao de %d", g.Points)
	} else {
		fmt.Printf("Você foi reprovado com a pontuacao de %d, estude um pouco mais", g.Points)
	}
}

func main() {
	game := &GameState{}

	ch := make(chan bool)
	csvLoadCh := make(chan bool)

	go game.ProcessCSV(ch, csvLoadCh)
	game.Init(ch)
	game.Run(csvLoadCh)
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("Não é permitido caractere diferente de numero, por favor insira um numero")
	}
	return i, nil
}
