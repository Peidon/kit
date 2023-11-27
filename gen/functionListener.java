// Generated from /Users/peidongxu/kit/ast/function.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link functionParser}.
 */
public interface functionListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link functionParser#f}.
	 * @param ctx the parse tree
	 */
	void enterF(functionParser.FContext ctx);
	/**
	 * Exit a parse tree produced by {@link functionParser#f}.
	 * @param ctx the parse tree
	 */
	void exitF(functionParser.FContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Function}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterFunction(functionParser.FunctionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Function}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitFunction(functionParser.FunctionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(functionParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(functionParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Decimal}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterDecimal(functionParser.DecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Decimal}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitDecimal(functionParser.DecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(functionParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(functionParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(functionParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(functionParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterString(functionParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link functionParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitString(functionParser.StringContext ctx);
}