package main

import (
	"context"
	"fmt"

	"time"
	"crypto/rand"
	"math/big"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Tai Xuan Jing tetragrams
// https://en.wiktionary.org/wiki/Taixuanjing
//
// There are 81 permutations of the following three lines, given 4 lines/grams
// ------ Heaven, --  -- Earth, - - - Man
// 3^4 = 81

type Passage struct {
    // id int // will be stored as a slice, so using index as id
    title string
    body string
}

func GetPassage(id int) (Passage) {
        var passages [82]Passage

	// passages[0] is unused to match id to index

	passages[1].title = "The Beginning of Power"
	passages[1].body = `The Tao that can be expressed<br>
	Is not the Tao of the Absolute.<br>
	The name that can be named<br>
	Is not the name of the Absolute.<br>
	<br>
	The nameless originated Heaven and Earth.<br>
	The named is the Mother of All Things.<br>
	<br>
	Thus, without expectation<br>
	One will always perceive the subtlety;<br>
	And, with expectation<br>
	One will always perceive the boundary.<br>
	<br>
	The source of these two is identical,<br>
	Yet their names are different.<br>
	Together they are called profound,<br>
	Profound and mysterious,<br>
	The gateway to the Collective Subtlety.`

	passages[2].title = "Using Polarity"
	passages[2].body =`When all the world knows beauty as beauty,<br>
	There is ugliness.<br>
	When they know good as good,<br>
	Then there is evil.<br>
	<br>
	In this way<br>
	Existence and nonexistence produce each other.<br>
	Difficult and easy complete each other.<br>
	Long and short contrast each other.<br>
	Pitch and tone harmonize each other.<br>
	Future and past follow each other.<br>
	<br>
	Therefore, Evolved Individuals<br>
	Hold their position without effort,<br>
	Practice their philosophy without words,<br>
	Are a part of All Things and overlook nothing.<br>
	They produce but do not possess,<br>
	Act wihtout expectation,<br>
	Succeed wtihout taking credit.<br>
	<br>
	Since, indeed, they take no credit, it remains with them.`

	passages[3].title = "Keeping Peace"
	passages[3].body =`And people will not contend.<br>
	Do not trasure goods that are hard to get,<br>
	And people will not become thieves.<br>
	Do not focus on desires,<br>
	And people's minds will not be confused.<br>
	<br>
	Therefore, Evolved Individuals lead others by<br>
	Opening their minds,<br>
	Reinforcing their centers,<br>
	Relaxing their desires,<br>
	Strengthening their characters.<br>
	<br>
	Let the people always act without strategy or desire;<br>
	Let the clever not venture to act.<br>
	Act without action,<br>
	And nothing is without order.`

	passages[4].title = "The Nature of the Tao"
	passages[4].body =`The Tao is empty and yet useful;<br>
	Somehow it never fills up.<br>
	So profound!<br>
	It resembles the source of All Things.<br>
	<br>
	It blunts the sharpness,<br>
	Unties the tangles,<br>
	And harmonizes the brightness.<br>
	It identifies with the ways of the world.<br>
	<br>
	So deep!<br>
	It resembles a certain existence.<br>
	I do not know whose offspring it is,<br>
	This Image in front of the source.`

	passages[5].title = "Holding to the Center"
	passages[5].body =`Heaven and Earth are impartial;<br>
	They regard All Things as straw dogs.<br>
	Evolved Individuals are impartial;<br>
	They regard all people as straw dogs.<br>
	<br>
	Between Heaven and Earth,<br>
	The space is like a bellows.<br>
	The shape changes,<br>
	But not the form.<br>
	The more it moves,<br>
	The more it produces.<br>
	<br>
	Too much talk will exhaust itself.<br>
	It is better to remain centered.`

	passages[6].title = "Perceiving the Subtle"
	passages[6].body =`The mystery of the valley is immortal;<br>
	It is known as the Subtle Female.<br>
	The gateway of the Subtle Female<br>
	Is the source of Heaven and Earth.<br>
	<br>
	Everlasting, endless, it appears to exist.<br>
	Its usefulness comes with no effort.`

	passages[7].title = "The Power of Selflessness"
	passages[7].body =`Heaven is eternal., the Earth everlasing.<br>
	They can be eternal and everlasting<br>
	Because they do not exist for themselves.<br>
	For that reason tehy can exist eternally.<br>
	<br>
	Therefore, Evolved Individuals<br>
	Put themselves last,<br>
	And yet they are first.<br>
	Put themselves outside,<br>
	And yet they remain.<br>
	<br>
	Is it not because they are without self-interest<br>
	That their interests succeed?`

	passages[8].title = "Noncompetitive Values"
	passages[8].body =`The highest value is like water.<br>
	<br>
	The value in water benefits All Things,<br>
	And yet it does not contend.<br>
	It stays in places that others disdain,<br>
	And therefore is close to the Tao.<br>
	<br>
	The value in a dwelling is location.<br>
	The value in a mind is depth.<br>
	The value in relations is benevolence.<br>
	The value in words is sincerity.<br>
	The value in leadership is order.<br>
	The value in work is competence.<br>
	The value in effort is timeliness.<br>
	<br>
	Since, indeed, they do not contend,<br>
	There is no resentment.`

	passages[9].title = "Transcending Decline"
	passages[9].body =`Holding to fullness<br>
	Is not as good as stopping in time.<br>
	<br>
	Sharpness that probes<br>
	Cannot protect for long.<br>
	<br>
	A house filled with riches<br>
	Cannot be defended.<br>
	<br>
	Pride in wealth and position<br>
	Is overlooking one's collapse.<br>
	<br>
	Withdrawing when success is achieved<br>
	Is the Tao in Nature.`

	passages[10].title = "Inner Harmony"
	passages[10].body =`In managing your instincts and embracing Oneness,<br>
	Can you be undivided?<br>
	In focusing your Influence,<br>
	Can you yield as a newborn child?<br>
	In clearing your insight,<br>
	Can you become free of error?<br>
	In loving people and leading the organization,<br>
	Can you take no action?<br>
	In opening and closing the gateway to nature,<br>
	Can you not weaken?<br>
	In seeing clearly in all directions,<br>
	Can you be wtihout knowledge?<br>
	<br>
	Produce things, cultivate things;<br>
	Produce but do not possess.<br>
	Act without expectation.<br>
	Advance without dominating.<br>
	These are called the Subtle Powers.`

	passages[11].title = "Using What Is Not"
	passages[11].body =`Thirty spokes converge at one hub;<br>
	What is not there makes the wheel useful.<br>
	Clay is shaped to form a vessel;<br>
	What is not there makes the vessel useful.<br>
	Doors and windows are cut to form a room;<br>
	What is not there makes the room useful.<br>
	<br>
	Therefore, take advantage of what is there,<br>
	By making use of what is not.`

	passages[12].title = "Controlling the Senses"
	passages[12].body =`The five colors will blind one's eye.<br>
	The five tones will deafen one's ear.<br>
	The five flavors will jade one's taste.<br>
	<br>
	Racing and hunting will drange one's mind.<br>
	Goods that are hard to get will obstruct one's way.<br>
	<br>
	Therefore, Evolved Individuals<br>
	Regard teh center and not the eye.<br>
	Hence they discard one and receive the other.`

	passages[13].title = "Expanding Identification"
	passages[13].body =`There is alarm in both favor and disgrace.<br>
	Esteem and fear are identified with the self.<br>
	<br>
	What is the meaning of "alarm in both favor and disgrace?"<br>
	Favor ascends; disgrace descends.<br>
	To attain them brings alarm.<br>
	To lose them brings alarm.<br>
	That is the meaning of "alarm in both favor and disgrace."<br>
	<br>
	What is the meaning of "esteem and fear are identified with the self?"<br>
	The reason for our fear<br>
	Is the presence of our self<br>
	When we are selfless,<br>
	What is there to fear?<br>
	<br>
	Therefore those who esteem the world as self<br>
	Will be committed to the world.<br>
	Those who love the world as self<br>
	Will be entrusted with the world.`

	passages[14].title = "The Essence of Tao"
	passages[14].body =`Looked at but not seen:<br>
	Its name is formless.<br>
	Listened to but not heard:<br>
	Its name is soundless.<br>
	Reached for but not obtained:<br>
	Its name is intangible.<br>
	<br>
	These three cannot be analyzed,<br>
	So they mingle and act as one.<br>
	<br>
	Its rising is not bright;<br>
	Its setting is not dark.<br>
	Endlessly, the nameless goes on,<br>
	Merging and returning to nothingness.<br>
	<br>
	That is why it is called<br>
	The form of the formless,<br>
	The image of nothingness.<br>
	That is why it is called elusive.<br>
	Confronted, its beginning is not seen.<br>
	Followed, its end is not seen.<br>
	<br>
	Hold on to the ancient Tao;<br>
	Control the current reality.<br>
	Be aware of the ancient origins;<br>
	This is called the Essence of Tao.`

	passages[15].title = "The Power in Subtle Force"
	passages[15].body =`Those skillful in teh ancient Tao<br>
	Are subtly ingenious and profoundly intuitive.<br>
	They are so deep they cannot be recognized.<br>
	Since, indeed, they cannot be recognized,<br>
	Their force can be contained.<br>
	<br>
	So careful!<br>
	As if wading a stream in winter.<br>
	So hesitant!<br>
	As if respecting all sides in the community.<br>
	So reserved!<br>
	As if acting as a guest.<br>
	So yielding!<br>
	As if ice about to melt.<br>
	So candid!<br>
	As if acting wiht simplicity.<br>
	So open!<br>
	As if acting as a valley.<br>
	So integrated!<br>
	As if acting as muddy water.<br>
	<br>
	Who can harmonize with muddy water,<br>
	And gradually arrive at clarity?<br>
	Who can move with stability,<br>
	And gradually bring endurance to life?<br>
	<br>
	Those who maintain the Tao<br>
	Do not desire to become full.<br>
	Indeed, since they are not full,<br>
	They can be used up and also renewed.`

	passages[16].title = "Knowing the Absolute"
	passages[16].body =`Attain the highest openness;<br>
	Maintain the deepest harmony.<br>
	Become a part of All Things;<br>
	In this way, I perceive the cycles.<br>
	<br>
	Indeed, things are numerous;<br>
	But each cycle merges with teh source.<br>
	Merging with the source is called harmonizing;<br>
	This is known as the cycle of destiny.<br>
	<br>
	The cycle of desitny is called the Absolute;<br>
	Knowing the Absolute is called insight.<br>
	To not know the Absolute<br>
	Is to recklessly become a part of misfortune.<br>
	<br>
	To know the Absolute is to be tolerant.<br>
	What is tolerant becomes impartial;<br>
	What is impartial becomes powerful;<br>
	What is powerful becomes natural;<br>
	What is natural becomes Tao.<br>
	<br>
	What has Tao becomes everlasting<br>
	And free from harm throughout life.`

	passages[17].title = "The Way of Subtle Influence"
	passages[17].body =`Superior leaders are those whose existence is merely known;<br>
	The next best are loved and honored;<br>
	The next are respected;<br>
	And the next are ridiculed.<br>
	<br>
	Those who lack belief<br>
	Will not in turn be believed.<br>
	But when the command comes from afar<br>
	And the work is done, the goal achieved,<br>
	The people say, "We did it naturally."`

	passages[18].title = "Losing the Instincts"
	passages[18].body =`When the great Tao is forgotten,<br>
	Philanthropy and morality appear.<br>
	Intelligent strategies are produced,<br>
	And great hypocrisies emerge.<br>
	<br>
	When the Family has no Harmony,<br>
	Piety and devotion appear.<br>
	The nation is confused by chaos,<br>
	And loyal patriots emerge.`

	passages[19].title = "Return to Simplicity"
	passages[19].body =`Discard the sacred, abandon strategies;<br>
	The people will benefit a hundredfold.<br>
	Discard philanthropy, abandon morality;<br>
	The people will return to natural love.<br>
	Discard cleverness, abandon the acquisitive;<br>
	The thieves will exist no longer.<br>
	<br>
	However, if these three passages are inadequate,<br>
	Adhere to these principles:<br>
	Perceive purity;<br>
	Embrace simplicity;<br>
	Reduce self-interest;<br>
	Limit desires.`

	passages[20].title = "Developing Independence"
	passages[20].body =`Discard the academic; have no anxiety.<br>
	How much difference is there between agrement and servility?<br>
	How much difference is there between good and evil?<br>
	That one should revere what others revere - how absurd and uncentered!<br>
	<br>
	The Collective Mind is expansive and flourishing,<br>
	As if receiving a great sacrifice,<br>
	As if ascending a living observatory.<br>
	<br>
	I alone remai uncommitted,<br>
	Like an infant who has not yet smiled,<br>
	Unattached, without a place to merge.<br>
	The Collective Mind is all-encompassing.<br>
	I alone seem to be overlooked.<br>
	I am unknowing to the core and unclear, unclear!<br>
	<br>
	Ordinary people are bright and obvious;<br>
	I alone am dark and obscure.<br>
	Ordinary people are exacting and sharp;<br>
	I alone am subdued and dull.<br>
	<br>
	Indifferent like the sea,<br>
	Ceaseless like a penetrating wind,<br>
	The Collective Mind is ever present.<br>
	And yet, I alone am unruly and remote.<br>
	I alone am different from the others<br>
	In treasuring nourishment from the Mother.`

	passages[21].title = "Knowing the Collective Origin"
	passages[21].body =`The natural expression of Power<br>
	Proceeds only through the Tao.<br>
	The Tao acts through Natural Law;<br>
	So formless, so intangible.<br>
	<br>
	Intangible, formless!<br>
	At its center appears teh Image.<br>
	Formless, intangible!<br>
	At its center appears Natural Law.<br>
	Obscure, mysterious!<br>
	At its center appears teh Life Force.<br>
	The Life Force is very real;<br>
	At its center appears truth.<br>
	<br>
	From ancient times to the present,<br>
	Its name ever remains,<br>
	Through the experience of the Collectie Origin.<br>
	<br>
	How do I know the way of the Collective Origin?<br>
	Through this.`

	passages[22].title = "Following the Pattern"
	passages[22].body =`What is curved becomes whole;<br>
	What is crooked becomes straight.<br>
	What is deep becomes filled;<br>
	What is exhausted becomes refreshed.<br>
	What is small becomes attainable;<br>
	What is excessive becomes confused.<br>
	<br>
	Thus Evolved Individuals hold to the One<br>
	And regard the world as their Pattern.<br>
	<br>
	They do not display themselves;<br>
	Therefore they are illuminated.<br>
	They do not define themselves;<br>
	Therefore they are distinguished.<br>
	They do not make claims;<br>
	Therefore they are credited.<br>
	They do not boast;<br>
	Therefore tehy advance.<br>
	<br>
	Since, indeed, they do not compete,<br>
	The world cannot compete with them.<br>
	<br>
	That ancient saying: "What is curved becomes whole" -<br>
	Are these empty words?<br>
	To become whole,<br>
	Turn within.`

	passages[23].title = "The Steady Force of Attitude"
	passages[23].body =`Nature rarely speaks.<br>
	Hence the whirlwind does not last a whole morning,<br>
	Nor the sudden rainstorm last a whole day.<br>
	What causes these?<br>
	Heaven and Earth.<br>
	If Heaven and Earth cannot make them long lasting,<br>
	How much less so can humans?<br>
	<br>
	Thus, those who cultivate the Tao<br>
	Identify with the Tao.<br>
	Those who cultivate Power<br>
	Identify with Power.<br>
	Those who cultivate failure<br>
	Identify with failure.<br>
	<br>
	Those who identify with the Tao<br>
	Are likewise welcomed by the Tao.<br>
	Those who identify with Power<br>
	Are likewise welcomed by Power.<br>
	Those who identify with failure<br>
	Are likewise welcomed by failure.<br>
	<br>
	Those who lack belief<br>
	Will not in turn be believed.`

	passages[24].title = "Nature rarely speaks"
	passages[24].body =`Those who are on tiptoe cannot stand firm.<br>
	Those who straddle cannot walk.<br>
	Those who display themselves cannot illuminate.<br>
	Those who define themselves cannot be distinguished.<br>
	Those who make claims can have no credit.<br>
	Those who boast cannot advance.<br>
	<br>
	To those who stay with the Tao,<br>
	These are like excess food and redundant actions<br>
	And are contrary to Natural Law.<br>
	Thus those who possess teh Tao turn away.`

	passages[25].title = "The Tao of Greatness"
	passages[25].body =`There was something in a state of fusion<br>
	Before Heaven and Earth were born.<br>
	<br>
	Silent, vast,<br>
	Independent, and unchanging;<br>
	Working everywhere, tirelessly;<br>
	It can be regarded as Mother of the world.<br>
	I do not know its name;<br>
	The word I say is Tao.<br>
	Forced to give it a name,<br>
	I say Great.<br>
	<br>
	Great means continuing.<br>
	Continuing means going far.<br>
	Going far means returning.<br>
	<br>
	Therefore the Tao is Great.<br>
	Heaven and Earth are Great.<br>
	A leader is likewise Great.<br>
	In the universe there are four Greatness,<br>
	And leadership is one of them.<br>
	<br>
	Humans are modeled on the earth.<br>
	The eart is modeled on heaven.<br>
	Heaven is modeled on the Tao.<br>
	The Tao is modeled on nature.`

	passages[26].title = "The Gravity of Power"
	passages[26].body =`Gravity is the foundation of levity.<br>
	Stillness is the master of agitation.<br>
	<br>
	Thus Evolved Individuals can travel the whole day<br>
	Without leaving behind their baggage.<br>
	However arresting the views,<br>
	They remain calm and unattached.<br>
	How can leaders with ten thousand chariots<br>
	Have a light-hearted position in the world?<br>
	<br>
	If they are light-hearted, they lose their foundation.<br>
	If they are agitated, they lose their mastery.`

	passages[27].title = "The Skillful Exchange of Information"
	passages[27].body =`A good path has no ruts.<br>
	A good speech has no flaws.<br>
	A good analysis uses no schemes.<br>
	<br>
	A good lock has no bar or bolt,<br>
	And yet it cannot be opened.<br>
	A good knot does not restrain,<br>
	And yet it cannot be unfastened.<br>
	<br>
	Thus Evolved Individuals are always good at saving others;<br>
	Hence no one is wasted.<br>
	They are always good at saving things;<br>
	Hence nothing is Wasted.<br>
	<br>
	This is called Doubling the LIght.<br>
	<br>
	Therefore a good person is teh teacher of an inferior person;<br>
	And an inferior person is teh resource of a good person.<br>
	One who does not treasure a teacher, or does not cherish a resource,<br>
	Although intelligent, is greatly deluded.<br>
	<br>
	This is called Significant Subtlety.`

	passages[28].title = "Uniting the Forces"
	passages[28].body =`Know the male,<br>
	Hold the female;<br>
	Become the world's stream.<br>
	By being teh world's streawm,<br>
	The Power will never leave.<br>
	This is returning to Infancy.<br>
	<br>
	Know the white,<br>
	Hold to the black;<br>
	Become the world's pattern.<br>
	By becoming the world's pattern,<br>
	The Power will never falter.<br>
	This is returning to Limitlessness.<br>
	<br>
	Know the glory,<br>
	Hold to the obscurity;<br>
	Become the world's valley.<br>
	By being the world's valley,<br>
	The Power will be sufficient.<br>
	This is returning to Simplicity.<br>
	<br>
	When Simplicity is broken up,<br>
	It is made into instruments.<br>
	Evolved Individuals who employ them,<br>
	Are made into leaders.<br>
	In this way the Great System is united.`

	passages[29].title = "The Way of Noninterference"
	passages[29].body =`Those who would take hold of the world and act on it,<br>
	Never, I notice, succeed.<br>
	<br>
	The world is a mysterious instrument,<br>
	Not made to be handled.<br>
	Those who act on it, spoil it.<br>
	Those who seize it, lose it.<br>
	<br>
	So, in Natural Law<br>
	Some lead, some follow;<br>
	Some agitate, some remain silent;<br>
	Some are firm, some are weak;<br>
	Some carry on, some lose heart.<br>
	<br>
	Thus, Evolved Individuals<br>
	Avoid extremes,<br>
	Avoid extravagance,<br>
	Avoid excess.`

	passages[30].title = "Leading the Leader"
	passages[30].body =`Those who use the Tao to guide leaders<br>
	Do not use forceful strategies in the world.<br>
	Such matters tend to recoil.<br>
	<br>
	Where armies are positioned,<br>
	Thorny brambles are produced.<br>
	A great military always brings years of hunger.<br>
	<br>
	Those who are skillful<br>
	Succeed and then stop.<br>
	They dare not hold on with force.<br>
	<br>
	They succeed and do not boast.<br>
	They succeed and do not make claims.<br>
	They succeed and are not proud.<br>
	They succeed and do not acquire in excess.<br>
	They succeed and do not force.<br>
	<br>
	Things overgrown will always decline.<br>
	This is not the Tao.<br>
	What is not the Tao will soon end.`

	passages[31].title = "The Use of Force"
	passages[31].body =`The finest weapons can be the instruments of misfortune,<br>
	And thus contrary to Natural Law.<br>
	Those who possess teh Tao turn away from them.<br>
	Evolved leaders occupy and honor the left;<br>
	Those who use weapons honor the right.<br>
	<br>
	Weapons are instruments of misfortune<br>
	That are used by the unevolved.<br>
	When their use is unavoidable,<br>
	The superior act with calm restraint.<br>
	<br>
	Even when victorious, let there be no joy,<br>
	For such joy leads to contentment with slaughter.<br>
	Those who are content with slaughter<br>
	Cannot find fulfillment in the world.`

	passages[32].title = "The Limits of Specialization"
	passages[32].body =`The Tao of the Absolute has no name.<br>
	Although infinitesimal in its Simplicity,<br>
	The world cannot master it.<br>
	<br>
	If leaders would hold on to it,<br>
	All Things would naturally follow.<br>
	Heaven and Earth would unite to rain Sweet Dew,<br>
	And people would naturally cooperate without commands.<br>
	<br>
	Names emerge when institutions begin.<br>
	When names emerge, know likewise to stop.<br>
	To know when to stop is to be free of danger.<br>
	<br>
	The presence of the Tao in the world<br>
	Is like the valley streawm joining the rivers and seas.`

	passages[33].title = "Self-Mastery"
	passages[33].body =`Those who know others are intelligent;<br>
	Those who know themselves have insight.<br>
	Those who master others have force;<br>
	Those who master themselves have strength.<br>
	<br>
	Those who know what is enough are wealthy.<br>
	Those who persevere have direction.<br>
	Those who maintain their position endure.<br>
	And those who die and yet do not perish, live on.`

	passages[34].title = "The Evolving Tao"
	passages[34].body =`The Great Tao extends everywhere.<br>
	It is on the left and the right.<br>
	<br>
	All Things depend on it for growth,<br>
	And it does not deny them.<br>
	It achieves its purpose,<br>
	And it does not have a name.<br>
	It clothes and cultivates All Things,<br>
	And it does not act as master.<br>
	<br>
	Always without desire,<br>
	It can be named Small.<br>
	All Things merge with it,<br>
	And it does not act as master.<br>
	It can be named Great.<br>
	<br>
	In the end it does not seek greatness,<br>
	And in that way the Great is achieved.`

	passages[35].title = "Sensing the Insensible"
	passages[35].body =`Hold fast to the Great Image,<br>
	And all the world will come.<br>
	Yet its coming brings no harm,<br>
	Only peace and order.<br>
	<br>
	When there is music together with food,<br>
	The audience will linger.<br>
	But when the Tao is expressed,<br>
	It seems without substance or flavor.<br>
	<br>
	We observe and there is nothing to see.<br>
	We listen and there is nothing to hear.<br>
	We use it and it is without end.`

	passages[36].title = "Concealing the Advantage"
	passages[36].body =`In order to deplete it,<br>
	It must be thoroughly extended.<br>
	In order to weaken it,<br>
	It must be thoroughly strengthened.<br>
	In order to reject it,<br>
	It must be thoroughly promoted.<br>
	In order to take away from it,<br>
	It must be thoroughly endowed.<br>
	<br>
	This is called a Subtle Insight.<br>
	The yielding can triumph over teh inflexible;<br>
	The weak can triumph over the strong.<br>
	Fish should not be taken from deep waters;<br>
	Nor should organizations make obvious their advantages.`

	passages[37].title = "The Power in Desirelessness"
	passages[37].body =`The Tao never acts,<br>
	And yet is never inactive.<br>
	<br>
	If leaders can hold on to it,<br>
	All Things will be naturally influenced.<br>
	Influenced and yet desiring to act,<br>
	I would calm them with Nameless Simplicity.<br>
	Nameless Simplicity is likewise without desire;<br>
	And without desire there is harmony.<br>
	<br>
	The world will then be naturally stabilized.`

	passages[38].title = "Power Without Motive"
	passages[38].body =`Superior Power is never Powerful, thus it has Power.<br>
	Inferior Power is always Powerful, thus it has no Power.<br>
	Superior Power takes no action and acts withotu motive.<br>
	Inferior Power takes action and acts with motive.<br>
	<br>
	Superior philanthropy takes action and acts without motive.<br>
	Superior morality takes action and acts with motive.<br>
	Superior propriety takes action and there is no response;<br>
	So it raises its arm to project itself.<br>
	<br>
	Therefore, lose the Tao and Power follows.<br>
	Lose the Power and philanthropy follows.<br>
	Lose philanthropy and morality follows.<br>
	Lose morality and propriety follows.<br>
	<br>
	One who has propriety has the veneer of truth<br>
	And yet is the leader of confusion.<br>
	One who knows the future has the luster of the Tao<br>
	And yet is ignorant of its origins.<br>
	<br>
	Therefore those wiht the greatest endurance<br>
	Can enter the substantial,<br>
	Not occupy its veneer;<br>
	Can enter reality,<br>
	Not occupy its luster.<br>
	Hence they discard one and receive the other.`

	passages[39].title = "Oneness in Leadership"
	passages[39].body =`From old, these may have harmony with the One:<br>
	<br>
	Heaven in harmony with the One becomes clear.<br>
	Earth in harmony with the One becomes stable.<br>
	Mind in harmony with the One becomes inspired.<br>
	Valleys in harmony wiht the One become full.<br>
	All Things in harmony with the One becomes creawtive.<br>
	Leaders in harmony with the One becomes incorruptible in the world.<br>
	<br>
	These were attained through Oneness.<br>
	<br>
	Heaven without clarity would probably crack.<br>
	Earth without stability would probably quake.<br>
	Mind without inspiration would probably sleep.<br>
	Valleys without fullness would probably dry up.<br>
	All Things without creativity would probably die off.<br>
	Leaders without incorruptible ways would probably stumble and fall.<br>
	<br>
	Indeed, the high-placed stem from teh humble;<br>
	The elevated are based upon the lowly.<br>
	This is why leaders call themselves<br>
	Alone, lonely, and unfavored.<br>
	Is this not because they stem form the humble and common?<br>
	Is it not?<br>
	<br>
	Therefore, attain honor without being honored.<br>
	Do not desire to shine like jade; wear ornaments as if they were stone.`

	passages[40].title = "The Way"
	passages[40].body =`Polarity is the movement of the Tao.<br>
	Receptivity is the way it is used.<br>
	The world and All Things were produced from its existence.<br>
	Its existence was produced from nonexistence.`

	passages[41].title = "Mastering the Paradox"
	passages[41].body =`When superior leaders hear of the Tao,<br>
	They diligently try to practice it.<br>
	When average leaders hear of the Tao,<br>
	They appear both aware and unaware of it.<br>
	When inferior leaders hear of the Tao,<br>
	They roar with laughter.<br>
	<br>
	Without sufficient laughter, it could not be the Tao;<br>
	Hence the long-established sayings:<br>
	<br>
	The Tao illuminated appears to be obscure;<br>
	The Tao advancing apperaas to be retreating;<br>
	The Tao leveled appears to be uneven.<br>
	<br>
	Superior Power appears to be low;<br>
	Great clarity appears to be spotted;<br>
	Extensive Power appears to be insufficient;<br>
	Established Power appers to be stolen;<br>
	Substantial Power appears to be spurious.<br>
	<br>
	The greatest space has no corners;<br>
	The greatest talents are slowly mastered;<br>
	The greatest music has the rarest sound;<br>
	The Great Image has no form.<br>
	<br>
	The Tao is hidden and nameless,<br>
	Yet it is the Tao that skillfully supports and completes.`

	passages[42].title = "Knowing Polarity"
	passages[42].body =`The Tao produced the One.<br>
	The One produced the Two.<br>
	The Two produced the Three.<br>
	The Three produced All Things.<br>
	<br>
	All Things carry Yin and hold to Yang;<br>
	Their blended Influence brings Harmony.<br>
	<br>
	People hate to be alone, lonely, and unfavored;<br>
	And yet leaders take these names.<br>
	<br>
	Thus in Natural Law<br>
	Some lose and in this way profit.<br>
	Some profit and in this way lose.<br>
	<br>
	What others have taught, I also teach:<br>
	Those who are violent do not die naturally.<br>
	I will make this my chief teaching.`

	passages[43].title = "Subtle Powers"
	passages[43].body =`The most yielding parts of the world<br>
	Overtake the most rigid parts of the world.<br>
	The insubstantial can penetrate continually.<br>
	<br>
	Therefore I know that wihtout action there is advantage.<br>
	<br>
	This philosophy wtihout words,<br>
	This advantage wtihout action -<br>
	It is rare, in the world, to attain them.`

	passages[44].title = "The Power in Needing Less"
	passages[44].body =`Which is dearer,<br>
	Name or life?<br>
	Which means more,<br>
	Life or wealth?<br>
	Which is worse,<br>
	Gain or loss?<br>
	<br>
	The stronger the attachments,<br>
	The greater the cost.<br>
	The more that is hoarded,<br>
	The deeper the loss.<br>
	<br>
	Know what is enough;<br>
	Be without disgrace.<br>
	Know when to stop;<br>
	Be without danger.<br>
	<br>
	In this way one lasts for a very long time.`

	passages[45].title = "Using Emptiness"
	passages[45].body =`If the greatest achievement is incomplete,<br>
	Then its usefulness is unimpaired.<br>
	If the greatest fullness is empty,<br>
	Then its usefulness is inexhaustible.<br>
	<br>
	The greatest directness is flexible.<br>
	The greatest skillfulness is awkward.<br>
	The greatest eloquence is hesitant.<br>
	<br>
	Agitation triumphs over the cold.<br>
	Stillness triumphs over the heated.<br>
	Clarity and stillness bring order to the world.`

	passages[46].title = "Knowing Enough"
	passages[46].body =`When the world possesses the Tao,<br>
	Even fast horses are used for their dung.<br>
	When the world is without the Tao,<br>
	War-horses are raised in the suburbs.<br>
	<br>
	There is no greater misfortune<br>
	Than not knowing what is enough.<br>
	There is no greater fault<br>
	Than desiring to acquire.<br>
	<br>
	Therefore know that enough is enough.<br>
	There will always be enough.`

	passages[47].title = "Cultivating Inner-Knowledge"
	passages[47].body =`Without going out of doors,<br>
	Know the world.<br>
	Without looking through the window,<br>
	See the Tao in Nature.<br>
	One may travel very far,<br>
	And know very little.<br>
	<br>
	Therefore, Evolved Individuals<br>
	Know without going about,<br>
	Recognize without looking,<br>
	Achieve without acting.`

	passages[48].title = "The Art of Nonaction"
	passages[48].body =`To pursue the academic, add to it daily.<br>
	To pursue the Tao, subtract from it daily.<br>
	Subtract and subtract again,<br>
	To arrive at nonaction.<br>
	Through nonaction nothing is left undone.<br>
	<br>
	The world is always held wtihout effort.<br>
	The moment there is effort,<br>
	The world is beyond holding.`

	passages[49].title = "Opening the Mind"
	passages[49].body =`Evolved Individuals have no fixed mind;<br>
	They make the mind of the People their mind.<br>
	<br>
	To those who are good, I am good;<br>
	To those who are not good, I am also good.<br>
	Goodness is Power.<br>
	<br>
	Of those who trust, I am trusting;<br>
	Of those who do not trust, I am also trusting.<br>
	Trust is Power.<br>
	<br>
	The Evolved Individuals in the world<br>
	Attract the world and merge with its mind.<br>
	The People all focus their eyes and ears;<br>
	Evolved Individuals all act as infants.`

	passages[50].title = "The Art of Survival"
	passages[50].body =`As life goes out, death comes in.<br>
	<br>
	Life has thirteen paths;<br>
	Death has thirteen paths.<br>
	Human life arrives at the realm of death<br>
	Also in thirteen moves.<br>
	<br>
	Why is this so?<br>
	Because life is lived lavishly.<br>
	<br>
	Now, as it is well known,<br>
	Those skilled in attracting life<br>
	Can travel across the land<br>
	And not meet a rhinoceros or tiger.<br>
	When the military comes in,<br>
	Their defence cannot be attacked.<br>
	<br>
	The rhinoceros is without a place to thrust its horn.<br>
	The tiger is without a place to affix its claw.<br>
	The military is without a place to admit its blade.<br>
	<br>
	Why is this so?<br>
	Because they are without the realm of death.`

	passages[51].title = "The Power of Impartial Support"
	passages[51].body =`The Tao produces;<br>
	Its Power supports;<br>
	Its Natural Law forms;<br>
	Its influence completes.<br>
	<br>
	Thus All Things without exception<br>
	Respect the Tao and value its Power.<br>
	To respect teh Tao and value its Power -<br>
	No one demands this, and it comes naturally.<br>
	<br>
	Therefore the Tao produces and its Power supports;<br>
	It advances, cultivates, comforts, matures, nourishes, and protects.<br>
	<br>
	Produce but do not possess.<br>
	Act without expectation.<br>
	Advance without dominating.<br>
	These are called the Subtle Powers.`

	passages[52].title = "Returning to Insight"
	passages[52].body =`The beginning of the world<br>
	May be regarded as the Mother of the world.<br>
	To apprehend the Mother,<br>
	Know the offspring.<br>
	To know the offspring<br>
	Is to remain close to the Mother,<br>
	And free from harm throughout life.<br>
	<br>
	Block the passages,<br>
	Close the doors;<br>
	In the end, life is idle.<br>
	<br>
	Open the pasages,<br>
	Increase undertakings;<br>
	In the end, life is hopeless.<br>
	<br>
	To perceive the small is called insight.<br>
	To remain yielding is called strength.<br>
	If, in using one's brightness,<br>
	One returns to insight,<br>
	Life will be free of misfortune.<br>
	<br>
	This is called learning the Absolute.`

	passages[53].title = "The Undivided Path"
	passages[53].body =`Using only a little knowledge,<br>
	I would travel the Great Way<br>
	And fear only of letting go.<br>
	The Great Way is very even;<br>
	Yet people love the byways.<br>
	<br>
	When an organization is divided,<br>
	Fields are overgrown,<br>
	Stores are empty,<br>
	Clothes are extravagant,<br>
	Sharp swords are worn,<br>
	Food and drink are excessive,<br>
	Wealth and treasure are hoarded.<br>
	<br>
	This is called stealing and exaggeration<br>
	And certainly not the Way!`

	passages[54].title = "Establishing a Universal View"
	passages[54].body =`What is skillfully established will not be uprooted;<br>
	What is skillfully grasped will not slip away.<br>
	Thus it is honored for generations.<br>
	<br>
	Cultivate the inner self;<br>
	Its Power becomes real.<br>
	Cultivate the home;<br>
	Its Power becomes abundant.<br>
	Cultivate the community;<br>
	Its Power becomes greater.<br>
	Cultivate the organization;<br>
	Its Power becomes prolific.<br>
	Cultivate the world;<br>
	Its Power becomes universal.<br>
	<br>
	Therefore through the inner self,<br>
	The inner self is conceived.<br>
	Through the home,<br>
	The home is conceived.<br>
	Through the community,<br>
	The community is conceived.<br>
	Through the organization,<br>
	The organization is conceived.<br>
	Through the world,<br>
	The world is conceived.<br>
	<br>
	How do I know the world?<br>
	Through this.`

	passages[55].title = "The Power in Not Contending"
	passages[55].body =`To possess Power that runs deep<br>
	Is to be like a newborn child.<br>
	<br>
	Poisonous insects do not sting it,<br>
	Fierce beasts do not seize it,<br>
	Birds of prey do not strike it.<br>
	<br>
	Its bones are yielding,<br>
	Its muscles are relaxed,<br>
	Its grip is strong.<br>
	<br>
	It does not yet know the union of male and female,<br>
	Yet its virility is active.<br>
	Its Life Force is at its greatest.<br>
	<br>
	It can scream all day,<br>
	Yet it does not become hoarse.<br>
	Its Harmony is at its greatest.<br>
	<br>
	To know Harmony is called the Absolute.<br>
	To know the Absolute is called insight.<br>
	To enhance life is called propitious.<br>
	To be conscious of Influence is called strength.<br>
	<br>
	Things overgrown must decline.<br>
	This is not the Tao.<br>
	What is not the Tao will soon end.`

	passages[56].title = "Gaining Oneness"
	passages[56].body =`Those who know do not speak.<br>
	Those who speak do not know.<br>
	<br>
	Block the passages.<br>
	Close the door.<br>
	Blunt the sharpness.<br>
	Untie the tangles.<br>
	Harmonize the brightness.<br>
	Identify wtih the ways of the world.<br>
	<br>
	This is called Profound Identification.<br>
	<br>
	It cannot be gained through attachment.<br>
	It cannot be gained through detachment.<br>
	It cannot be gained through advantage.<br>
	It cannot be gained through disadvantage.<br>
	It cannot be gained through esteem.<br>
	It cannot be gained through humility.<br>
	<br>
	Hence it is the treasure of the world.`

	passages[57].title = "The Power in Effortlessness"
	passages[57].body =`Lead the organization with correctness.<br>
	Direct the military with surprise tactics.<br>
	Take hold of the world with effortlessness.<br>
	<br>
	How do I know it is so?<br>
	Through this:<br>
	<br>
	Too many prohibitions in the world,<br>
	And people become insufficient.<br>
	Too many sharp weapons among people,<br>
	And the nation grows confused.<br>
	Too much cunning strategy among people,<br>
	And strange things start to happen.<br>
	Too obvious a growth in laws and regulations,<br>
	And too many criminals emerge.<br>
	<br>
	Thus Evolved Individuals say:<br>
	<br>
	Look to nonaction,<br>
	And people will be naturally influenced.<br>
	Look to refined tranquillity,<br>
	And peopel will be naturally correct.<br>
	Look to effortlessness,<br>
	And people will be naturally affluent.<br>
	Look to nondesire,<br>
	And people will be naturally simple.`

	passages[58].title = "Cultivating the Center"
	passages[58].body =`If the administration is subdued,<br>
	The people are sincere.<br>
	If the administration is exacting,<br>
	The people are deficient.<br>
	<br>
	Misfortune! Good fortune supports it.<br>
	Good Fortune! Misfortune hides within.<br>
	Who knows where it ends?<br>
	Is there no order?<br>
	<br>
	Order can revert to the unusual;<br>
	Good can revert to teh abnormal;<br>
	And people indeed are bewildered<br>
	For a long, long time.<br>
	<br>
	Thus Evolved INdividuals are<br>
	Square without dividing;<br>
	Honest without offending;<br>
	Straightforward without straining;<br>
	Bright without dazzling.`

	passages[59].title = "The Way of Moderation"
	passages[59].body =`In leading people and serving Nature,<br>
	There is nothing better than moderation.<br>
	Since, indeed, moderation means yielding early;<br>
	Yielding early means accumulating Power.<br>
	<br>
	When Power is accumulated,<br>
	Nothing is impossible.<br>
	When nothing is impossible,<br>
	One knows no limits.<br>
	One who knows no limits<br>
	Can possess the organization.<br>
	<br>
	An organization that possesses the Mother<br>
	Can endure and advance.<br>
	This means deep roots and firm foundation:<br>
	Durability and longevity through observation of the Tao.`

	passages[60].title = "Holding the Position"
	passages[60].body =`Leading a large organization is like cooking a small fish.<br>
	<br>
	If the Tao is present in the world,<br>
	The cunning are not mysterious.<br>
	Not only are the cunning not mysterious,<br>
	Their mystery does not harm others.<br>
	<br>
	Not only does their mystery not harm others,<br>
	The Evolved also do not harm others.<br>
	Since together they do no harm,<br>
	The Power returns and accumulates.`

	passages[61].title = "The Power in Modesty"
	passages[61].body =`A large organization should flow downward<br>
	To intersect with the world.<br>
	It is the female of the world.<br>
	The female always overcomes the male by stillness;<br>
	Through stillness, she makes herself low.<br>
	<br>
	Thus if a large organization<br>
	Is lower than a small organization,<br>
	It can receive the small organization.<br>
	And if the small organization<br>
	Stays lower than a large organization,<br>
	It can receive the large organization.<br>
	<br>
	Therefore one receives by becoming low;<br>
	Another receives by being low.<br>
	<br>
	Yet what a large organization desires<br>
	Is to unite and support others.<br>
	And what a small organization desires<br>
	Is to join and serve others.<br>
	<br>
	So for both to gain the position they desire,<br>
	The larger should place itself low.`

	passages[62].title = "The Tao in Leaders"
	passages[62].body =`The Tao is a refuge for All Things,<br>
	The treasure of the good,<br>
	The protector of the not good.<br>
	<br>
	Honor can be bought with fine words;<br>
	Others can be joined with fine conduct.<br>
	So if some are not good,<br>
	Why waste them?<br>
	<br>
	In this way the Emperor is established;<br>
	The three officials are installed.<br>
	And although the large jade disc<br>
	Is preceded by a team of horses,<br>
	This is not as good as sitting,<br>
	Advancing in the Tao.<br>
	<br>
	Why did those of old treasure the Tao?<br>
	Did they not say:<br>
	Seek it and it is attained;<br>
	Posses faults and they are released?<br>
	Thus it is the treasure of the world.`

	passages[63].title = "The Path of Least Resistance"
	passages[63].body =`Act without action; work without effort.<br>
	Taste without savoring.<br>
	Magnify the small; increase the few.<br>
	Repay ill-will with kindness.<br>
	<br>
	Plan the difficult when it is easy;<br>
	Handle the big where it is small.<br>
	The world's hardest work begins when it is easy;<br>
	The world's largest effort begins where it is small.<br>
	Evolved Individuals, finally, take no great action,<br>
	And in that way the great is achieved.<br>
	<br>
	Those who commit easily, inspire little trust,<br>
	How easy to inspire hardness!<br>
	Therefore Evolved Individuals view all as difficult.<br>
	Finally they have no difficulty!`

	passages[64].title = "The Power at the Beginning"
	passages[64].body =`What is at rest is easy to hold;<br>
	What is not yet begun is easy to plan.<br>
	What is thin is easy to melt;<br>
	What is minute is easy to disperse.<br>
	Deal with things before they emerge;<br>
	Put them in order before there is disorder.<br>
	<br>
	A tree of many arm spans is produced from a tiny sprout.<br>
	A tower of nine stories is raised from a pile of earth.<br>
	A journey of a thousand miles begin with a footstep.<br>
	<br>
	Those who act on things, spoil them;<br>
	Those who seize things, lose them.<br>
	Thus Evolved Individuals do nothing;<br>
	Hence they spoil nothing.<br>
	They seize nothing;<br>
	Hence they lose nothing.<br>
	<br>
	People often spoil their work at the point of its completion.<br>
	With care at the end as well as the beginning,<br>
	No work will be spoiled.<br>
	<br>
	Thus Evolved Individuals desire to be desireless<br>
	And do not treasure goods that are hard to get.<br>
	They learn without learning,<br>
	By returning to the place where the Collective Mind passes.<br>
	In this way they assist All Things naturally<br>
	Without venturing to act.`

	passages[65].title = "The Danger in Cleverness"
	passages[65].body =`Those skillful in the ancient Tao<br>
	Are not obvious to the people.<br>
	They appear to be simple-minded.<br>
	<br>
	People are difficult to lead<br>
	Because they are too clever.<br>
	Hence, to lead the organization with cleverness<br>
	Will harm the organization.<br>
	To lead the organization without cleverness<br>
	Will benefit the organization.<br>
	<br>
	Those who know these two things<br>
	Have investigated the patterns of the Absolute.<br>
	To know and investigate the patterns<br>
	Is called the Subtle Power.<br>
	<br>
	The Subtle Power is profound and far-reaching.<br>
	Together wtih the Natural Law of polarity,<br>
	It leads to the Great Harmony.`

	passages[66].title = "The Power in Staying Low"
	passages[66].body =`The rivers and seas lead the hundred streams<br>
	Because they are skillful at staying low.<br>
	Thus they are able to lead the hundren streams.<br>
	<br>
	Therefore, to rise above people,<br>
	One must, in speaking, stay below them.<br>
	To remain in front of people,<br>
	One must put onself behind them.<br>
	<br>
	Therefore Evolved Individuals remain above,<br>
	And yet the people are not weighted down.<br>
	They remain in front,<br>
	And the people are not held back.<br>
	<br>
	Therefore the world willingly elects them,<br>
	And yet it does not reject them.<br>
	Because they do not compete,<br>
	The world cannot compete with them.`

	passages[67].title = "The Power in Compassion"
	passages[67].body =`All the world thinks that my Tao is great;<br>
	And yet it seems inconceivable.<br>
	Only its greatness makes it seem inconceivable.<br>
	If it could be conceived of,<br>
	It would have become insignificant long ago.<br>
	<br>
	I have Three Treasures that support and protect;<br>
	The first is compassion.<br>
	The second is moderation.<br>
	The third is daring not to be first in the world.<br>
	<br>
	With compassion one becomes courageous;<br>
	With moderation one becomes expansive.<br>
	In daring not to be first in the world,<br>
	One becomes the instrument of leadership.<br>
	<br>
	Now if one is courageous without compassion,<br>
	Or expansive without moderation,<br>
	Or first without holding back,<br>
	One is doomed!<br>
	<br>
	Compassion always triumphs when attacked;<br>
	It brings security when maintained.<br>
	Nature aids its leaders<br>
	By arming them with compassion.`

	passages[68].title = "Nonaggressive Strength"
	passages[68].body =`A skillful leader does not use force.<br>
	A skillful fighter does not feel anger.<br>
	A skillful master does not engage the opponent.<br>
	A skillful employer remains low.<br>
	<br>
	This is called the power in not contending.<br>
	This is called the strength to employ others.<br>
	This is called the highest emulation of Nature.`

	passages[69].title = "Neutralizing Escalation"
	passages[69].body =`The strategists have a saying:<br>
	"I dare not act as a host,<br>
	Yet I act as a guest,<br>
	I dare not advance an inch,<br>
	Yet I retreat a foot."<br>
	<br>
	This is called<br>
	Traveling without moving,<br>
	Rising up without arms,<br>
	Projecting without resistance,<br>
	Capturing without strategies.<br>
	<br>
	No misfortune is greater than underestimating resistance;<br>
	Underestimating resistance will destroy my Treasures.<br>
	Thus when mutually opposing strategies escalate,<br>
	The one who feels sorrow will triumph.`

	passages[70].title = "Knowing the Tao"
	passages[70].body =`My words are very easy to know,<br>
	Very easy to follow.<br>
	Yet the world is unable to know them,<br>
	Unable to follow them.<br>
	<br>
	My words have a source,<br>
	My efforts have mastery.<br>
	Indeed, since none know this,<br>
	They do not know me.<br>
	The rare ones who know me<br>
	Must treasure me.<br>
	<br>
	Therefore, Evolved Individuals<br>
	Wear a coarse cloth covering<br>
	With precious jade at the center.`

	passages[71].title = "Knowing the Disease"
	passages[71].body =`To know that you do not know is best.<br>
	To not know of knowing is a disease.<br>
	<br>
	Indeed, to be sick of the disease,<br>
	Is the way to be free of the disease.<br>
	<br>
	Evolved Individuals are free of the disease,<br>
	Because they are sick of the disease.<br>
	<br>
	This is the way to be free of disease.`

	passages[72].title = "The Appropriate Perspective "
	passages[72].body =`If the people do not fear authority,<br>
	Then authority will expand.<br>
	Do not disrespect their position;<br>
	Do not reject their lives.<br>
	Since indeed they are not rejected,<br>
	They do not reject.<br>
	<br>
	Therefore, enlightened people know themselves,<br>
	But do not display themselves.<br>
	They love themselves,<br>
	But do not treasure themselves.<br>
	<br>
	Hence they discard one and receive the other.`

	passages[73].title = "Nature's Way"
	passages[73].body =`Those bold in daring will die;<br>
	Those bold in not daring will survive.<br>
	Of these two, either may benefit or harm.<br>
	<br>
	Nature decides which is evil,<br>
	But who can know why?<br>
	Even Evolved Individuals regard this as difficult.<br>
	<br>
	The Tao in Nature<br>
	Does not contend,<br>
	Yet skillfully triumphs.<br>
	Does not speak,<br>
	Yet skillfully responds.<br>
	Does not summon,<br>
	And yet attracts.<br>
	Does not hasten,<br>
	Yet skillfully designs.<br>
	<br>
	Nature's network is vast, so vast,<br>
	Its mesh is coarse, yet nothing slips through.`

	passages[74].title = "Unnatural Authority"
	passages[74].body =`When people do not fear death,<br>
	How can they be threatened with death?<br>
	Suppose people fear death and still do not conform,<br>
	Who would dare seize them and put them to death?<br>
	<br>
	There is always the Master Executioner who kills.<br>
	To substitute for the Master Executioner in killing<br>
	Is like substituting for the Master Carpenter who carves.<br>
	Whoever substitutes for the Master Carpenter in carving,<br>
	Rarely escapes injury to his hands.`

	passages[75].title = "Self-Destructive Leadership"
	passages[75].body =`People are hungry.<br>
	Because those above consume too much in taxes,<br>
	People are hungry.<br>
	<br>
	People are difficult to lead.<br>
	Because those above interfere with them,<br>
	People are difficult to lead.<br>
	<br>
	People make light of death.<br>
	Because those above deeply seek survival,<br>
	People make light of death.<br>
	<br>
	Indeed, it is those who do not interfere with life<br>
	Who are capable of respecting life.`

	passages[76].title = "The Power in Flexibility"
	passages[76].body =`A man living is yielding and receptive.<br>
	Dying, he is rigid and inflexible.<br>
	All Things, the grass and trees:<br>
	Living, they are yielding and fragile;<br>
	Dying, they are dry and withered.<br>
	<br>
	Thus those who are firm and inflexible<br>
	Are in harmony with dying.<br>
	Those who are yielding and receptive<br>
	Are in harmony with living.<br>
	<br>
	Therefore an inflexible strategy will not triumph;<br>
	An inflexible tree will be attacked.<br>
	The position of the highly inflexible will descend;<br>
	The position of the yielding and receptive will ascend.`

	passages[77].title = "Directing the Power"
	passages[77].body =`The Tao in Nature<br>
	Is like a bow that is stretched.<br>
	The top is pulled down,<br>
	The bottom is raised up.<br>
	What is excessive is reduced,<br>
	What is insufficient is supplemented.<br>
	<br>
	The Tao in Nature<br>
	Reduces the excessive<br>
	And supplements the insufficient.<br>
	The Tao in Man is not so;<br>
	He reduces the insufficient,<br>
	Because he serves the excessive.<br>
	<br>
	Who then can use excess to serve the world?<br>
	Those who possess the Tao.<br>
	<br>
	Therefore Evolved Individuals<br>
	Act without expectation,<br>
	Succeed without taking credit,<br>
	And have no desire to display their excellence.`

	passages[78].title = "Accepting the Blame"
	passages[78].body =`Nothing in the world,<br>
	Is as yielding and receptive as water;<br>
	Yet in attacking the firm and inflexible,<br>
	Nothing triumphs so well.<br>
	Because of what it is not,<br>
	This becomes easy.<br>
	<br>
	The receptive triumphs over the inflexible;<br>
	The yielding triumphs over the rigid.<br>
	None in the world do not know this.<br>
	None have the ability to practice it.<br>
	<br>
	Therefore Evolved Individuals say:<br>
	One who accepts the disgrace of the organization<br>
	Can be called the leader of the grain shrine.<br>
	One who accepts the misfortunes of the organization<br>
	Can be called the leader of the world.<br>
	<br>
	Right words appear to reverse themselves.`

	passages[79].title = "The Power In Not Taking Advantage"
	passages[79].body =`Even when a great resentment is reconciled,<br>
	Some resentment must linger,<br>
	How can this be made good?<br>
	<br>
	That is why Evolved Individuals<br>
	Hold the left side of the contract<br>
	And do not censure others.<br>
	Those with Power are in charge of the contract;<br>
	Those without Power are in charge of resolving it.<br>
	<br>
	The Tao in Nature has no favorites.<br>
	It always works through the good person.`

	passages[80].title = "Fulfilling Independence"
	passages[80].body =`In a small organization with a few people;<br>
	<br>
	Let there be ten or a hundred times<br>
	More tools than they can use.<br>
	Let the people value their lives<br>
	And yet not move far away.<br>
	Even though there are boats and carriages,<br>
	There is no occasion to use them.<br>
	Even though there are armor and weapons,<br>
	There is no occasion to display them.<br>
	<br>
	Let the people again knot cords and use them.<br>
	Their food will be pleasing.<br>
	Their clothes will be fine.<br>
	Their homes will be secure.<br>
	Their customs will be joyful.<br>
	<br>
	Nearby organizations may watch each other;<br>
	Their crowing and barking may be heard.<br>
	Yet the people may grow old and die<br>
	Without coming or going between them.`

	passages[81].title = "The Evolved Way"
	passages[81].body =`Sincere words are not embellished;<br>
	Embellished words are not sincere.<br>
	Those who are good are not defensive;<br>
	Those who are defensive are not good.<br>
	Those who know are not erudite;<br>
	Those who are erudite do not know.<br>
	<br>
	Evolved Individuals do not accumulate.<br>
	The more they do for others, the more they gain;<br>
	The more they give to others, the more they possess.<br>
	<br>
	The Tao of Nature<br>
	Is to serve without spoiling.<br>
	The Tao of Evolved Individuals<br>
	Is to act wihtout contending.`

	return passages[id]
}

func GenerateId() int {
	base10_sum := 0 // This will be the base10 sum of the monograms

	// Simple random
	// choose random from [0,81)
	// randBigInt, err := rand.Int(rand.Reader, big.NewInt(81))
	// if err != nil {
	// 	panic(err)
	// }

	// More authentic random to build the tetragram (tetra is four)
	// Choose a random int 0-2 for each "gram" of the tetragram, then sum it to the previous with a weighting
	// To convert from trinary/ternary (base 3) to base 10, multiply each part by it's order of magnitude
	// i.e. The first or bottom line gets a weight of 27, the second, 9, then 3, then 1
	// TODO: Am I doing this backwards? Would it be more authentic to give the last rand trinary the heaviest weight?
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randBigInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		randInt := int(randBigInt.Int64())

		switch i {
			case 0:
				base10_sum = base10_sum + randInt*27
			case 1:
				base10_sum = base10_sum + randInt*9
			case 2:
				base10_sum = base10_sum + randInt*3
			case 3:
				base10_sum = base10_sum + randInt*1
		}
	}

	return base10_sum
}

func GenerateTetragram(id int) string {
	tetragrams := []string{"𝌆", "𝌇", "𝌈", "𝌉", "𝌊", "𝌋", "𝌌", "𝌍", "𝌎", "𝌏", "𝌐", "𝌑", "𝌒", "𝌓", "𝌔", "𝌕", "𝌖", "𝌗", "𝌘", "𝌙", "𝌚", "𝌛", "𝌜", "𝌝", "𝌞", "𝌟", "𝌠", "𝌡", "𝌢", "𝌣", "𝌤", "𝌥", "𝌦", "𝌧", "𝌨", "𝌩", "𝌪", "𝌫", "𝌬", "𝌭", "𝌮", "𝌯", "𝌰", "𝌱", "𝌲", "𝌳", "𝌴", "𝌵", "𝌶", "𝌷", "𝌸", "𝌹", "𝌺", "𝌻", "𝌼", "𝌽", "𝌾", "𝌿", "𝍀", "𝍁", "𝍂", "𝍃", "𝍄", "𝍅", "𝍆", "𝍇", "𝍈", "𝍉", "𝍊", "𝍋", "𝍌", "𝍍", "𝍎", "𝍏", "𝍐", "𝍑", "𝍒", "𝍓", "𝍔", "𝍕", "𝍖"}

	return tetragrams[id]
}

func GenerateHTML(id int) string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
        // No longer linking to text, including it
        // href := fmt.Sprintf("<a href='https://terebess.hu/english/tao/Wing.html#Kap%02d'>%s</a>", id, id)

	tetragram := GenerateTetragram(id)
	passage := GetPassage(id)
	html := fmt.Sprintf(`<html>
<head><style>
/* === BASE HEADING === */
h1 { position: relative; padding: 0; margin: 0; font-family: "Raleway", sans-serif; font-weight: 300; font-size: 40px; color: #080808; -webkit-transition: all 0.4s ease 0s; -o-transition: all 0.4s ease 0s; transition: all 0.4s ease 0s; }
h1 span { display: block; font-size: 0.5em; line-height: 1.3; }
h1 em { font-style: normal; font-weight: 600; }

/* === HEADING STYLE === */
.title h1 { text-transform: capitalize; }
.title h1:before { position: absolute; left: 0; bottom: 0; width: 60px; height: 2px; content: ""; background-color: #c50000; }
.title h1 span { font-size: 13px; font-weight: 500; text-transform: uppercase; letter-spacing: 4px; line-height: 3em; padding-left: 0.25em; color: rgba(0, 0, 0, 0.4); padding-bottom: 10px; }
.footer { font-size: 11px; font-weight: 500; text-transform: uppercase; line-height: 3em; padding-left: 0.25em; color: rgba(0, 0, 0, 0.4); padding-bottom: 10px; }
</style></head>
<body>
<div class="title"><h1>%02d %s
<span>%s</span>
</h1></div>
<br>
%s
<br><br>
<span class="footer">%s</span>
</body>
</html>`, id, tetragram, passage.title, passage.body, now)
	return html
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := GenerateId()
	body := GenerateHTML(id)
	// fmt.Println("This message will show up in the CLI console.")
	// Headers:         map[string]string{"Content-Type": "text/plain; charset=utf-8"},

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Body:            body,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
