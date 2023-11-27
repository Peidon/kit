// Generated from /Users/peidongxu/kit/ast/function.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class functionLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, MUL=4, DIV=5, ADD=6, SUB=7, IDENTIFIER=8, DECIMAL=9, 
		NUMBER=10, WHITESPACE=11, STRING=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "MUL", "DIV", "ADD", "SUB", "IDENTIFIER", "DECIMAL", 
			"NUMBER", "WHITESPACE", "STRING", "ESC", "UNICODE", "SAFECODEPOINT", 
			"HEX"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "','", "')'", "'*'", "'/'", "'+'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, "MUL", "DIV", "ADD", "SUB", "IDENTIFIER", "DECIMAL", 
			"NUMBER", "WHITESPACE", "STRING"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public functionLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "function.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\fu\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0005"+
		"\u00072\b\u0007\n\u0007\f\u00075\t\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007:\b\u0007\n\u0007\f\u0007=\t\u0007\u0005\u0007?\b\u0007"+
		"\n\u0007\f\u0007B\t\u0007\u0001\b\u0004\bE\b\b\u000b\b\f\bF\u0001\b\u0001"+
		"\b\u0004\bK\b\b\u000b\b\f\bL\u0003\bO\b\b\u0001\t\u0004\tR\b\t\u000b\t"+
		"\f\tS\u0001\n\u0004\nW\b\n\u000b\n\f\nX\u0001\n\u0001\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0005\u000b`\b\u000b\n\u000b\f\u000bc\t\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0003\fj\b\f\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f"+
		"\u0001\u000f\u0000\u0000\u0010\u0001\u0001\u0003\u0002\u0005\u0003\u0007"+
		"\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b"+
		"\u0017\f\u0019\u0000\u001b\u0000\u001d\u0000\u001f\u0000\u0001\u0000\u0007"+
		"\u0003\u0000AZ__az\u0004\u000009AZ__az\u0001\u000009\u0003\u0000\t\n\r"+
		"\r  \b\u0000\"\"//\\\\bbffnnrrtt\u0003\u0000\u0000\u001f\"\"\\\\\u0003"+
		"\u000009AFaf{\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000"+
		"\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000"+
		"\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000"+
		"\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000"+
		"\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000"+
		"\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000"+
		"\u0001!\u0001\u0000\u0000\u0000\u0003#\u0001\u0000\u0000\u0000\u0005%"+
		"\u0001\u0000\u0000\u0000\u0007\'\u0001\u0000\u0000\u0000\t)\u0001\u0000"+
		"\u0000\u0000\u000b+\u0001\u0000\u0000\u0000\r-\u0001\u0000\u0000\u0000"+
		"\u000f/\u0001\u0000\u0000\u0000\u0011D\u0001\u0000\u0000\u0000\u0013Q"+
		"\u0001\u0000\u0000\u0000\u0015V\u0001\u0000\u0000\u0000\u0017\\\u0001"+
		"\u0000\u0000\u0000\u0019f\u0001\u0000\u0000\u0000\u001bk\u0001\u0000\u0000"+
		"\u0000\u001dq\u0001\u0000\u0000\u0000\u001fs\u0001\u0000\u0000\u0000!"+
		"\"\u0005(\u0000\u0000\"\u0002\u0001\u0000\u0000\u0000#$\u0005,\u0000\u0000"+
		"$\u0004\u0001\u0000\u0000\u0000%&\u0005)\u0000\u0000&\u0006\u0001\u0000"+
		"\u0000\u0000\'(\u0005*\u0000\u0000(\b\u0001\u0000\u0000\u0000)*\u0005"+
		"/\u0000\u0000*\n\u0001\u0000\u0000\u0000+,\u0005+\u0000\u0000,\f\u0001"+
		"\u0000\u0000\u0000-.\u0005-\u0000\u0000.\u000e\u0001\u0000\u0000\u0000"+
		"/3\u0007\u0000\u0000\u000002\u0007\u0001\u0000\u000010\u0001\u0000\u0000"+
		"\u000025\u0001\u0000\u0000\u000031\u0001\u0000\u0000\u000034\u0001\u0000"+
		"\u0000\u00004@\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u000067\u0005"+
		".\u0000\u00007;\u0007\u0000\u0000\u00008:\u0007\u0001\u0000\u000098\u0001"+
		"\u0000\u0000\u0000:=\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000"+
		";<\u0001\u0000\u0000\u0000<?\u0001\u0000\u0000\u0000=;\u0001\u0000\u0000"+
		"\u0000>6\u0001\u0000\u0000\u0000?B\u0001\u0000\u0000\u0000@>\u0001\u0000"+
		"\u0000\u0000@A\u0001\u0000\u0000\u0000A\u0010\u0001\u0000\u0000\u0000"+
		"B@\u0001\u0000\u0000\u0000CE\u0007\u0002\u0000\u0000DC\u0001\u0000\u0000"+
		"\u0000EF\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000FG\u0001\u0000"+
		"\u0000\u0000GN\u0001\u0000\u0000\u0000HJ\u0005.\u0000\u0000IK\u0007\u0002"+
		"\u0000\u0000JI\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LJ\u0001"+
		"\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000MO\u0001\u0000\u0000\u0000"+
		"NH\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000\u0000O\u0012\u0001\u0000"+
		"\u0000\u0000PR\u0007\u0002\u0000\u0000QP\u0001\u0000\u0000\u0000RS\u0001"+
		"\u0000\u0000\u0000SQ\u0001\u0000\u0000\u0000ST\u0001\u0000\u0000\u0000"+
		"T\u0014\u0001\u0000\u0000\u0000UW\u0007\u0003\u0000\u0000VU\u0001\u0000"+
		"\u0000\u0000WX\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XY\u0001"+
		"\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000Z[\u0006\n\u0000\u0000[\u0016"+
		"\u0001\u0000\u0000\u0000\\a\u0005\"\u0000\u0000]`\u0003\u0019\f\u0000"+
		"^`\u0003\u001d\u000e\u0000_]\u0001\u0000\u0000\u0000_^\u0001\u0000\u0000"+
		"\u0000`c\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000ab\u0001\u0000"+
		"\u0000\u0000bd\u0001\u0000\u0000\u0000ca\u0001\u0000\u0000\u0000de\u0005"+
		"\"\u0000\u0000e\u0018\u0001\u0000\u0000\u0000fi\u0005\\\u0000\u0000gj"+
		"\u0007\u0004\u0000\u0000hj\u0003\u001b\r\u0000ig\u0001\u0000\u0000\u0000"+
		"ih\u0001\u0000\u0000\u0000j\u001a\u0001\u0000\u0000\u0000kl\u0005u\u0000"+
		"\u0000lm\u0003\u001f\u000f\u0000mn\u0003\u001f\u000f\u0000no\u0003\u001f"+
		"\u000f\u0000op\u0003\u001f\u000f\u0000p\u001c\u0001\u0000\u0000\u0000"+
		"qr\b\u0005\u0000\u0000r\u001e\u0001\u0000\u0000\u0000st\u0007\u0006\u0000"+
		"\u0000t \u0001\u0000\u0000\u0000\f\u00003;@FLNSX_ai\u0001\u0006\u0000"+
		"\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}