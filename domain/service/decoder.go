package service

import (
	"errors"
	"ml-challenge/app/api"
	"ml-challenge/util"
	"strings"
)

type DecoderService struct{}

func NewDecoderService() *DecoderService {
	return &DecoderService{}
}

func (s *DecoderService) getLocation(m1, m2, m3 api.MessageApi) (x, y float32, e error) {
	/*
		La distancia d entre dos puntos P1(x1, y1) y P2(x2,y2) se relacionan entre si con la siguiente ecuacion
		d^2 = x1^2 - 2x1x2 + x2^2 + y1^2 - 2y1y2 + y2^2
		por lo tanto tenemos 2 incognitas con un sistema de 3 ecuaciones que llamaremos ec1, ec2, ec3
		Igualamos a cero las ecuaciones y tratamos de reducirlas entre si para despejar x e y:
		ec2 - ec1 = 0 y ec3 - ec1 = 0 nos permite despejar x e y para formar un sistema lineal homogeneo de 2 ecuaciones
		y 2 incognitas que podemos resolver con la regla de cramer.
	*/
	s1 := m1.Receiver
	s2 := m2.Receiver
	s3 := m3.Receiver

	return util.Cramer([2][3]float32{
		{
			2 * (s2.PosX - s1.PosX),
			2 * (s2.PosY - s1.PosY),
			calculateDelta(m1, m2),
		},
		{
			2 * (s3.PosX - s1.PosX),
			2 * (s3.PosY - s1.PosY),
			calculateDelta(m1, m3),
		},
	})
}

func (s *DecoderService) getMessage(m1, m2, m3 api.MessageApi) (string, error) {
	if !(len(m1.Message) == len(m2.Message) && len(m2.Message) == len(m3.Message)) {
		// como no estaba especificado en el challente asumo que los mensajes tiene todos el mismo largo y que la
		// perdida de info se representa por un string vacio.
		return "", errors.New("the messages are of different length")
	}
	result := make([]string, 0, len(m1.Message))
	for i := 0; i < len(m1.Message); i++ {
		/*
		 * En este caso el algoritmo no tiene en cuenta si en la misma posicion en los 3 mensajes
		 * los strings que no son vacios tienen algun tipo de prioridad entre si, como no fue pedido en el challenge no
		 * agregue nada.
		 * Ej:
		 * m1 = ["hola"], m2 = ["alo"], m3 = [""]
		 * resultado = "hola" solo por como esta implementado
		 */

		aux := util.CompareStringsNotEmpty(m3.Message[i], util.CompareStringsNotEmpty(m1.Message[i], m2.Message[i]))
		if aux == "" {
			continue
		}
		result = append(result, aux)
	}
	return strings.Join(result, " "), nil
}

func (s *DecoderService) Decode(m1, m2, m3 api.MessageApi) (*api.DecodedMessageApi, error) {
	x, y, err := s.getLocation(m1, m2, m3)
	if err != nil {
		return nil, errors.New("we could not find a matching coordinate for the message")
	}
	m, err := s.getMessage(m1, m2, m3)
	if err != nil {
		return nil, errors.New("we could not decode the message")
	}
	return &api.DecodedMessageApi{
		Pos:     api.PositionApi{X: x, Y: y},
		Message: m,
	}, err
}

func calculateDelta(m1, m2 api.MessageApi) float32 {
	// Funcion para calcular el delta de la resta de ecuaciones de distancia
	deltaX := util.Float32Sqrt(m2.Receiver.PosX) - util.Float32Sqrt(m1.Receiver.PosX)
	deltaY := util.Float32Sqrt(m2.Receiver.PosY) - util.Float32Sqrt(m1.Receiver.PosY)
	deltaDistance := util.Float32Sqrt(m1.Distance) - util.Float32Sqrt(m2.Distance)
	return deltaX + deltaY + deltaDistance
}
