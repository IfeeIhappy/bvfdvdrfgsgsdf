package peer
import "github.com/gskartwii/rbxfile"

type Packet83_01 struct {
	Instance *rbxfile.Instance
}

func DecodePacket83_01(packet *UDPPacket, context *CommunicationContext) (interface{}, error) {
	var err error
	inner := &Packet83_01{}
	thisBitstream := packet.Stream
    referent, err := thisBitstream.ReadObject(false, context)
    inner.Instance = context.InstancesByReferent.TryGetInstance(referent)

	return inner, err
}